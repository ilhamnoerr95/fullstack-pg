package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Backend running on :8080")

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Backend"))
	}))
}