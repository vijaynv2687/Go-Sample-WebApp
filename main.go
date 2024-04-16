package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<h1>Welcome to my site built in golang!</h1>")
	tmpl := template.Must(template.ParseFiles("Registration.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", nil)
}
