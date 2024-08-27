package main

import (
	"context"
	"context-app/customers"
	"context-app/products"
	"fmt"
	"log"
	"time"

	"net/http"

	_ "net/http/pprof"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func LoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s - %s", r.Method, r.URL.Path)
		start := time.Now()
		handler.ServeHTTP(w, r)
		// statusCode := w.
		elapsed := time.Since(start)
		log.Printf("%s %v\n", msg, elapsed)
	})
}

func TraceIdMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.New()
		valCtx := context.WithValue(r.Context(), "request-id", reqId)
		handler.ServeHTTP(w, r.WithContext(valCtx))
	})
}

func main() {
	// log.SetFormatter(&log.JSONFormatter{})

	r := mux.NewRouter()
	r.Use(LoggerMiddleware)
	r.Use(TraceIdMiddleware)

	// r.HandleFunc("/", index.Handler)
	r.HandleFunc("/customers", customers.Handler)
	r.HandleFunc("/products", products.GetAllHandler).Methods("GET")
	r.HandleFunc("/products", products.AddProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", products.GetAProductHandler).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
		return
	}
}
