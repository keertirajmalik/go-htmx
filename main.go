package main

import (
	"log"
	"net/http"
)

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@test.com"),
			newContact("Clara", "cd@test.com"),
		},
	}
}

func handleContactGet(templates *Templates, data Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.Render(w, "contact", data)
	}
}

func main() {
	const port = "8080"

	templates := newTemplates()
	count := &Count{Count: 0}
    data := newData()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /count", handleCountGet(templates, count))
	mux.HandleFunc("POST /count", handleCountInc(templates, count))

	mux.HandleFunc("GET /contact", handleContactGet(templates, data))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
