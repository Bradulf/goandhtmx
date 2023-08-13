package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
				{Title: "Schindler's List", Director: "Steven Spielberg"},
				{Title: "Raging Bull", Director: "Martin Scorsese"},
				{Title: "Casablanca", Director: "Michael Curtiz"},
			},
		}

		tmplt.Execute(w, films)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
