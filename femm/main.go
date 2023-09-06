package main

import (
	"fmt"
	"net/http"
	"text/template"

	"frontendmasters.com/go/museum/api"
	"frontendmasters.com/go/museum/data"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hellos from Go program"))
		},
	)
	server.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("./templates/index.tmpl")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}

		html.Execute(w, data.GetAll())

	})
	server.HandleFunc("/api/exibitions", api.Get)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while opening the server")
	}
}
