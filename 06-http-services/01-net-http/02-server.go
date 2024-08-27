package main

import (
	"encoding/json"
	"fmt"
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
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlerFn := appServer.routes[r.URL.Path]; handlerFn != nil {
		handlerFn(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

func (appServer *AppServer) Register(pattern string, handlerFn func(http.ResponseWriter, *http.Request)) {
	appServer.routes[pattern] = handlerFn
}

// application specific implementation
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers data will be served!")
}

func main() {
	appServer := NewAppServer()
	appServer.Register("/", IndexHandler)
	appServer.Register("/products", ProductsHandler)
	appServer.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
