package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	templates := newTemplates()
	count := &Count{Count: 0}

    page := newPage()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /count", handleCountGet(templates, count))
	mux.HandleFunc("POST /count", handleCountInc(templates, count))

	mux.HandleFunc("GET /contact", handleContactGet(templates, page))
	mux.HandleFunc("POST /contact", handleContactCreate(templates, page))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
