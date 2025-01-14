package main

import (
	"fmt"
	"net/http"
	"text/template"

	ascii "ascii/functions"
)

func main() {
	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/form", formFunc)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func formFunc(w http.ResponseWriter, r *http.Request) {
	tp2, _ := template.ParseFiles("index.html")
	tp2.Execute(w, nil)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	tp1, _ := template.ParseFiles("home.html")

	worr := r.FormValue("word")
	typee := r.FormValue("typee")

	laste := ascii.Ascii(worr , typee)

	tp1.Execute(w, laste) 
}
