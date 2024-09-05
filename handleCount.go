package main

import (
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

func handleCountGet(templates *Templates, count *Count) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.Render(w, "count", count)
	}
}

func handleCountInc(templates *Templates, count *Count) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		templates.Render(w, "count-div", count)
	}
}
