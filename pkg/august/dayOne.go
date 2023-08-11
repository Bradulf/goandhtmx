package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("index.html"))
		tmplt.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
