package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/customers", func(r chi.Router) {
		r.Get("/", getAllCustomers)
		r.Post("/", createCustomer)
		r.Get("/{id}", getCustomer)
		r.Put("/{id}", updateCustomer)
		r.Delete("/{id}", deleteCustomer)
	})

	http.ListenAndServe(":3000", r)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getAllCustomers"))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createCustomer"))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getCustomer"))
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updateCustomer"))
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleteCustomer"))
}
