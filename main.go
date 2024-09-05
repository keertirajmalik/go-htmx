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


    fs := http.FileServer(http.Dir("css"))
    mux.Handle("/css/", http.StripPrefix("/css/", fs))

    fs = http.FileServer(http.Dir("images"))
    mux.Handle("/images/", http.StripPrefix("/images/", fs))

	mux.HandleFunc("GET /count", handleCountGet(templates, count))
	mux.HandleFunc("POST /count", handleCountInc(templates, count))

	mux.HandleFunc("GET /contact", handleContactGet(templates, page))
	mux.HandleFunc("POST /contact", handleContactCreate(templates, page))
	mux.HandleFunc("DELETE /contact/{id}", handleContactDelete(templates, page))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
