package main

import (
	"net/http"
	"strconv"
	"time"
)

type Contact struct {
	Name  string
	Email string
	Id    int
}

var id = 0

func newContact(name, email string) Contact {
	id++
	return Contact{
		Name:  name,
		Email: email,
		Id:    id,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (d *Data) indexOf(id int) int {

	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}

	return -1
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@test.com"),
			newContact("Clara", "cd@test.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func handleContactGet(templates *Templates, page Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.Render(w, "contact", page)
	}
}

func handleContactCreate(templates *Templates, page Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			w.WriteHeader(http.StatusUnprocessableEntity)
			templates.Render(w, "contact-form", formData)
			return
		}

		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		templates.Render(w, "contact-form", newFormData())
		templates.Render(w, "oob-contact-div", contact)
	}
}

func handleContactDelete(_ *Templates, page Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid id"))
			return
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Contact not found"))
			return
		}

        time.Sleep(3 * time.Second)

		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)

		w.WriteHeader(http.StatusNoContent)
	}
}
