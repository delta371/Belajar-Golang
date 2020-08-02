package handler

import (
	"fmt"
	"net/http"
	"text/template"
	"web/model"

)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Welcome to my page")
	// fmt.Fprintln(w, "This is my firs page")

	t, err := template.ParseFiles("./html/template_index.html")
	if err != nil {
		fmt.Println(err)
	}

	cus := model.Customer{"Muhamad Ridwan", 22, "Programmer Backend"}

	t.Execute(w, cus)

}
