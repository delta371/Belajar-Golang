// ! Kumpulan dari statmen yang di jalankan dengan bersamaan dalam 1 task
// ! dibagi 2 yaitu function yang mengembalikan value dengan function tan tidak mengembalikan value

package main

import "fmt"

func main() {
	var a string = "Muhamad Ridwan"
	fmt.Println(a)
	Cetaknama()
	fmt.Println(CetakValue())
	CetaknamaMemakaiParameter("Ridwan")
	fmt.Println(CetakValueParameterReturn("Bogor"))
	Pertambahan(2, 7)

}

func Cetaknama() {
	var Job = "Saya adalah seorang Backend development"
	fmt.Println(Job)
}

func CetakValue() string { //jika ingin me return maka deklarasi tipe datanya
	var Hobby = "Bermain Komputer"
	return Hobby
}

func CetaknamaMemakaiParameter(nama string) {
	fmt.Println(nama)
}

func CetakValueParameterReturn(lahir string) string {
	return lahir
}

func cetakNamaPrivate() {
	var nama = "Beben Solar"
	fmt.Println(nama)
}

func Pertambahan(a, b int) {
	fmt.Println("Hasil pertambahan: ", a+b)
}

// * private --> Kalo Huruf depannya kecil
// * public ---> Kalo huruf depannya Besar
