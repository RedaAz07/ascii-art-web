package main

import (
	"fmt"
	"net/http"
	"text/template"
	"unicode"

	ascii "ascii/functions"
)

func main() {
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))))

	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/", formFunc)

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func formFunc(w http.ResponseWriter, r *http.Request) {
	tp2, _ := template.ParseFiles("index.html")
	if r.Method != "GET" {
		http.Error(w, "Bad Request - GET Only", 405)
		return
	}




	tp2.Execute(w, nil)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	tp1, _ := template.ParseFiles("home.html")

	worr := r.FormValue("word")
	typee := r.FormValue("typee")

	if worr == "" || typee == "" {
		http.Error(w, "Word and Type are required", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(worr); i++ {
		if unicode.IsLetter(rune(worr[i])) && (worr[i] < 32 || worr[i] > 126) {
			http.Error(w, " there is a special charts in ur  text ..... ", 400)
			return
		}
	}

	laste := ascii.Ascii(worr, typee,w)
	if r.Method != "POST" {
		http.Error(w, "Bad Request - POST Only", http.StatusBadRequest)
		return
	}
	tp1.Execute(w, laste)
}
