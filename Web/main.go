// ! di golang memiliki package net/http
// ! tujuannya untuk pembuatan rooting web server,blackpink dll

package main

import (
	"net/http"
	"web/handler"

)

func main() {
	http.HandleFunc("/", handler.Index)           //Menentukkan index
	http.HandleFunc("/barang", handler.GetBarang) //Menentukkan index

	http.ListenAndServe(":5000", nil) //Membuat server
}
