package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hellos from Go program"))
		},
	)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while opening the server")
	}
}
