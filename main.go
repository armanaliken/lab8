package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name   string
	Height string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.ParseFiles("static/index.html")
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			t.Execute(w, nil)
		} else if r.Method == "POST" {
			name := r.FormValue("name")
			height := r.FormValue("height")

			person := Person{
				Name:   name,
				Height: height,
			}

			t, err := template.ParseFiles("static/result.html")
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			t.Execute(w, person)
		}
	})

	http.ListenAndServe(":8080", nil)
}
