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
        if r.URL.Path == "/styles/" {
            http.Redirect(w, r, "/notfound", http.StatusFound)
            return
        }
        http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))).ServeHTTP(w, r)
    }))

    http.HandleFunc("/ascii-art", ResultFunc)
    http.HandleFunc("/", formFunc)
    http.HandleFunc("/notfound", notFoundFunc)
    http.HandleFunc("/MethodNotAllowed", MethodNotAllowedFunc)

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
        http.Redirect(w, r, "/MethodNotAllowed", http.StatusFound)
        return
    }

    tp2.Execute(w, nil)
}

func ResultFunc(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/ascii-art" {
        http.Redirect(w, r, "/notfound", http.StatusFound)
        return
    }

    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/MethodNotAllowed", http.StatusFound)
        return
    }

    tp2, _ := template.ParseFiles("template/index.html")

    word := r.FormValue("word")
    typee := r.FormValue("typee")

    var errorMessage string
    if word == "" {
        errorMessage = "Please enter a word."
    } else if typee == "" {
        errorMessage = "Please select a type."
    } else if len(word) > 5000 {
        errorMessage = "The word length should not exceed 5000 characters."
    } else {
        for i := 0; i < len(word); i++ {
            if !unicode.IsPrint(rune(word[i])) {
                errorMessage = "Invalid character in the word."
                break
            }
        }
    }

    if errorMessage != "" {
        data := struct {
            Error string
        }{
            Error: errorMessage,
        }
        tp2.Execute(w, data)
        return
    }

    tp1, _ := template.ParseFiles("template/result.html")

    laste := ascii.Ascii(word, typee, w)
    if laste == "" {
        http.Redirect(w, r, "/badrequest", http.StatusFound)
        return
    }

    tp1.Execute(w, laste)
}

func notFoundFunc(w http.ResponseWriter, r *http.Request) {
    tp, _ := template.ParseFiles("template/notfound.html")
    w.WriteHeader(http.StatusNotFound)
    tp.Execute(w, nil)
}

func MethodNotAllowedFunc(w http.ResponseWriter, r *http.Request) {
    tp, _ := template.ParseFiles("template/MethodNotAllowed.html")
    w.WriteHeader(http.StatusMethodNotAllowed)
    tp.Execute(w, nil)
}
