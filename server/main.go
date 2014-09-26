package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Jokes API!"))
	})

	http.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("insert joke here"))
	})

	http.ListenAndServe(":8080", nil)
}
