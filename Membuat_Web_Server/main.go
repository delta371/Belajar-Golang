package main

import (
	"fmt"
	"net/http"

)

func formData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Halaman Tidak Di Temukan", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "formBooking.html")

	case "POST":
		var kesalahan = r.ParseForm()
		if kesalahan != nil {
			fmt.Println(w, "Ada Kesalahan : ", kesalahan)
			return
		}

		var kode = r.FormValue("kode")
		var nama = r.FormValue("nama")
		var lapangan = r.FormValue("lapangan")
		fmt.Fprintln(w, "KODE = ", kode)
		fmt.Fprintln(w, "NAMA = ", nama)
		fmt.Fprintln(w, "LAPANGAN = ", lapangan)

	default:
		fmt.Fprintln(w, "Maaf, Method yang didukung hanya GET dan POST")
	}
}

func main() {
	http.HandleFunc("/", formData)

	fmt.Println("Web berjalan di server localhost:5000")
	http.ListenAndServe(":5000", nil)

}
