package main

import (
	"log"
	"net/http"
	"sync"
)


type Count struct {
	Count int
	mu    sync.Mutex
}

func (c *Count) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Count++
}

func handleGet(templates *Templates, count *Count) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.Render(w, "index", count)
	}
}

func handleCountInc(templates *Templates, count *Count) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		templates.Render(w, "count", count)
	}
}

func main() {
	const port = "8080"

	templates := newTemplates()
	count := &Count{Count: 0}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleGet(templates, count))
	mux.HandleFunc("POST /count", handleCountInc(templates, count))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
