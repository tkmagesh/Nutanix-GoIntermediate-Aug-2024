package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

// dummy data
var products = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
	{103, "Marker", 50},
}

type AppServer struct {
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "error parsing the payload", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers data will be served!")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
