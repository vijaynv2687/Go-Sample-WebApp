package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Registration.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on three thousand:3000")
	http.ListenAndServe(":3000", nil)
}
