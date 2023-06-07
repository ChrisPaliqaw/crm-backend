package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type serial uint64

type Customer struct {
	Id        serial `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var customerDatabase = map[serial]Customer{
	1: {
		1,
		"Logan Roy",
		"CEO",
		"logan@waystar.com",
		"212-111-1111",
		false,
	},
	2: {
		2,
		"Tom Wambsgans",
		"Head of Media",
		"twambsgans@waystar.com",
		"212-222-2222",
		true,
	},
	3: {
		3,
		"Nan Pierce",
		"CEO",
		"nan@pgm.com",
		"212-333-3333",
		false,
	},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/customers", func(r chi.Router) {
		r.Get("/", getCustomers)
		r.Post("/", addCustomer)
		r.Get("/{id}", getCustomer)
		r.Put("/{id}", updateCustomer)
		r.Delete("/{id}", deleteCustomer)
	})

	http.ListenAndServe(":3000", r)
}

func writeJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(v)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("getAllCustomers"))

	writeJson(w, customerDatabase)
}

func returnError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		returnError(w)
		return
	}
	c, ok := customerDatabase[serial(id)]
	if !ok {
		returnError(w)
		return
	}
	writeJson(w, c)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("createCustomer"))

	// req.Header.Set("User-Agent", "spacecount-tutorial")
	if r.Body == nil {
		returnError(w)
		return
	}

	defer r.Body.Close()

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
		returnError(w)
		return
	}

	newCustomer := Customer{}
	jsonErr := json.Unmarshal(body, &newCustomer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		returnError(w)
		return
	}

	_, ok := customerDatabase[newCustomer.Id]
	if ok {
		returnError(w)
		return
	}
	customerDatabase[newCustomer.Id] = newCustomer
	writeJson(w, newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		returnError(w)
		return
	}

	defer r.Body.Close()

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
		returnError(w)
		return
	}

	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		returnError(w)
		return
	}

	newCustomer := Customer{}
	jsonErr := json.Unmarshal(body, &newCustomer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		returnError(w)
		return
	}

	if id != uint64(newCustomer.Id) {
		returnError(w)
		return
	}

	_, ok := customerDatabase[newCustomer.Id]
	if !ok {
		returnError(w)
		return
	}
	customerDatabase[newCustomer.Id] = newCustomer
	writeJson(w, newCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("deleteCustomer"))

	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		returnError(w)
		return
	}

	customerToDelete, ok := customerDatabase[serial(id)]
	if !ok {
		returnError(w)
		return
	}
	delete(customerDatabase, serial(id))
	writeJson(w, customerToDelete)
}
