package main

import (
	"encoding/json"
	"net/http"

)

type Barang struct {
	Nama        string `json:"nama"`
	Harga       int    `json:"harga"`
	JenisBarang string `json:"jenis_barang"`
}

var data = []Barang{
	{"AAA", 670000, "Makanan"},
	{"BBBB", 70000, "Minuman"},
	{"CCCC", 80000, "Makanan"},
	{"DDDD", 90000, "Minuman"},
}

func main() {
	http.HandleFunc("/barang", GetBarang)
	http.HandleFunc("/barang-id", GetBarangByNama)

	http.ListenAndServe(":8000", nil)
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenct-type", "aplication/json") //untuk menentukan type response nya sebagai json

	if r.Method == "POST" {
		res, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return // jika error maka mereturn error
		}

		//jika tidak ada error maka kirimkan datanya sebagai response
		w.Write(res)
		return
	}

	http.Error(w, "400 status bad request", http.StatusBadRequest) //jika ada yan memakai post maka akan error
}

func GetBarangByNama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenct-type", "aplication/json")

	if r.Method == "GET" {
		nama := r.FormValue("nama")

		for _, v := range data {
			if v.Nama == nama {
				result, err := json.Marshal(v)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
	}
	http.Error(w, "400 status bad request", http.StatusBadRequest)
}
