package main

import (
	"log"
	"net/http"
)

func main() {
	initDb()

	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler)
	log.Println("Server runnning on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
