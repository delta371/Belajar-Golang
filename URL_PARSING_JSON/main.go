package main

import (
	// "fmt"
	// "net/url"
	// "strconv"
	"encoding/json"
	"fmt"
)

type Barang struct {
	Nama        string `json:"nama"`
	Harga       int    `json:"harga"`
	JenisBarang string `json:"jenis_barang"`
}

func main() {
	// a := 4.85
	// fmt.Println(a)
	// fmt.Println(int(a))

	// b := "900"
	// c, _ := strconv.Atoi(b) //mengubah string ke int
	// fmt.Println(c + 400)

	// var d = "http://www.belajar-golang.com/basic?name=Ridwan&tinggi=160"
	// fmt.Println(d)

	// e, _ := url.Parse(d) //parsing si d nya ke URL

	// fmt.Println(e.Scheme)  //http
	// fmt.Println(e.Host)    // www.belajar-golang.com
	// fmt.Println(e.Path)    // /basic
	// fmt.Println(e.Query()) // map[name:[Ridwan] tinggi:[160]]

	// name := e.Query()["name"][0] // [0]mengambil nma tanpa aa array
	// tinggi := e.Query()["tinggi"][0]
	// fmt.Println(name)
	// fmt.Println(tinggi)

	DecodeJSONToObject()
	DecodeArrayJSONToObject()
	DecodeJSONToMapStringInterface()
	DecodeJSONToInterface()
	EncodeJSONObject()
}

func DecodeJSONToObject() {
	jsonString := `{"nama": "AAAA", "harga": 7800, "jenis_barang": "Makanan" }`
	jsonData := []byte(jsonString)

	var Bar Barang

	err := json.Unmarshal(jsonData, &Bar) // memasukkan json ke bar , bar nya harus memiliki tanda (&)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Bar.Nama)
	fmt.Println(Bar.Harga)
	fmt.Println(Bar.JenisBarang)
}

func DecodeArrayJSONToObject() {
	jsonString := `[
		{"nama": "AAAA", "harga": 7800, "jenis_barang": "Makanan" },
		{"nama": "BBBB", "harga": 8800, "jenis_barang": "Minuman" },
		{"nama": "CCCC", "harga": 9200, "jenis_barang": "Makanan" },
		{"nama": "DDDD", "harga": 9500, "jenis_barang": "Minuman" }
	]`
	jsonData := []byte(jsonString)

	var Bar []Barang

	err := json.Unmarshal(jsonData, &Bar) // memasukkan json ke bar , bar nya harus memiliki tanda (&)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range Bar {
		// fmt.Println(v)
		fmt.Println("Nama Barang :", v.Nama)
		fmt.Println("Harga Barang :", v.Harga)
	}

}

func DecodeJSONToMapStringInterface() {
	jsonString := `{"nama": "AAAA", "harga": 7800, "jenis_barang": "Makanan" }`
	jsonData := []byte(jsonString)

	var Bar map[string]interface{}

	err := json.Unmarshal(jsonData, &Bar) // memasukkan json ke bar , bar nya harus memiliki tanda (&)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Bar["jenis_barang"]) //jika memakai map maka memanggilnya dengan key
	fmt.Println(Bar["harga"])        //jika memakai map maka memanggilnya dengan key
}

func DecodeJSONToInterface() {
	jsonString := `{"nama": "AAAA", "harga": 7800, "jenis_barang": "Makanan" }`
	jsonData := []byte(jsonString)

	var Bar interface{}

	err := json.Unmarshal(jsonData, &Bar) // memasukkan json ke bar , bar nya harus memiliki tanda (&)
	if err != nil {
		fmt.Println(err)
	}

	decode := Bar.(map[string]interface{}) //mengubah Bar yang tadi nya interface menjadi map seperti di atas
	fmt.Println(decode["nama"])
}

func EncodeJSONObject() {
	data := Barang{"AAA", 78000, "Makanan"}
	res, _ := json.Marshal(data) // mengencode json, hasilnya berbentuk byte

	fmt.Println(string(res)) //mengconvert dari byte ke string
}
