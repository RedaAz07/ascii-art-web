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

	http.HandleFunc("/ascii-art", ResultFunc)
	http.HandleFunc("/", formFunc)

	http.HandleFunc("/notfound", notFoundFunc)
	// Default handler for non-existent routes
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/notfound", http.StatusFound)
	})

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func formFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	tp2, _ := template.ParseFiles("template/index.html")
	if r.Method != "GET" {
		http.Error(w, "Bad Request - GET Only", http.StatusMethodNotAllowed)
		return
	}

	tp2.Execute(w, nil)
}

func ResultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	tp1, _ := template.ParseFiles("template/result.html")

	word := r.FormValue("word")
	typee := r.FormValue("typee")



	for i := 0; i < len(word); i++ {
		if unicode.IsLetter(rune(word[i])) && (word[i] < 32 || word[i] > 126) {
			http.Error(w, "There are special characters in your text.", http.StatusBadRequest)
			return
		}
	}

	laste := ascii.Ascii(word, typee, w)
	if r.Method != "POST" {
		http.Error(w, "Bad Request - POST Only", http.StatusMethodNotAllowed)
		return
	}

	tp1.Execute(w, laste)
}

func notFoundFunc(w http.ResponseWriter, r *http.Request) {
	tp, _ := template.ParseFiles("template/notfound.html")
	w.WriteHeader(http.StatusNotFound)
	tp.Execute(w, nil)
}
