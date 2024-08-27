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

// Library
type AppServer struct {
	routes      map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)
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
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		middleware := appServer.middlewares[i]
		handlerFn = middleware(handlerFn)
	}
	appServer.routes[pattern] = handlerFn
}

func (appServer *AppServer) UseMiddleware(middleware func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)) {
	appServer.middlewares = append(appServer.middlewares, middleware)
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

// middlewares
func logMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next(w, r)
	}
}

func JSONMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func main() {

	appServer := NewAppServer()
	appServer.UseMiddleware(JSONMiddleware)
	appServer.UseMiddleware(logMiddleware)
	appServer.Register("/", IndexHandler)
	appServer.Register("/products", ProductsHandler)
	appServer.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
