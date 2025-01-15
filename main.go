package main

import (
	"fmt"
	"net/http"
	"text/template"
	"unicode"

	ascii "ascii/functions"
)

func main() {
	http.Handle("/styles/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/styles/" || r.URL.Path == "/styles" {
			http.Redirect(w,r,"/notfound",303)
			return
		}
		http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))).ServeHTTP(w, r)
	}))

	http.HandleFunc("/ascii-art", ResultFunc)
	http.HandleFunc("/", formFunc)
	http.HandleFunc("/notfound", notFoundFunc)

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
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request - GET Only", http.StatusMethodNotAllowed)
		return
	}

	tp2.Execute(w, nil)
}

// Result page handler
func ResultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	// Redirect if the request is not coming from the form
	if r.Referer() != "http://localhost:8080/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request - POST Only", http.StatusMethodNotAllowed)
		return
	}

	// Parse the result template
	tp1, _ := template.ParseFiles("template/result.html")

	// Get form values
	word := r.FormValue("word")
	typee := r.FormValue("typee")

	// Validate input
	if word == "" || typee == "" {
		http.Error(w, "Bad Request - Please fill out the form", http.StatusBadRequest)
		return
	}

	// Check for invalid characters
	for i := 0; i < len(word); i++ {
		if unicode.IsLetter(rune(word[i])) && (word[i] < 32 || word[i] > 126) {
			http.Error(w, "Bad Request - Invalid characters in your text", http.StatusBadRequest)
			return
		}
	}

	// Generate ASCII art
	laste := ascii.Ascii(word, typee, w)
	if laste == "" {
		http.Error(w, "Internal Server Error - Failed to generate ASCII art", http.StatusInternalServerError)
		return
	}

	// Execute the result template with the generated ASCII art
	tp1.Execute(w, laste)
}

// 404 Not Found handler
func notFoundFunc(w http.ResponseWriter, r *http.Request) {
	tp, _ := template.ParseFiles("template/notfound.html")
	w.WriteHeader(http.StatusNotFound)
	tp.Execute(w, nil)
}
