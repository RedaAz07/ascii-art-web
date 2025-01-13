package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	css:=http.StripPrefix("/styles/",http.FileServer(http.Dir("styles")))
	http.Handle("/styles/",css)
	http.HandleFunc("/", home)
	fmt.Println("Server running at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("index.html")
	temp.Execute(w, nil)
}
