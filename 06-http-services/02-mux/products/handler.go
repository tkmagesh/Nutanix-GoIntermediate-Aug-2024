package products

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"req-id": r.Context().Value("request-id"),
	}).Info("[Products Handler]")

	ps := NewProductsService()

	if err := json.NewEncoder(w).Encode(ps.GetAll(r.Context())); err != nil {
		http.Error(w, "error encoding data", http.StatusInternalServerError)
	}

}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"req-id": r.Context().Value("request-id"),
	}).Info("[Products Handler]")

	ps := NewProductsService()
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
	}
	ps.AddNew(r.Context(), newProduct)
	w.WriteHeader(http.StatusCreated)

}

func GetAProductHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"req-id": r.Context().Value("request-id"),
	}).Info("[Products Handler]")

	ps := NewProductsService()
	var vars = mux.Vars(r)
	if pid, err := strconv.Atoi(vars["id"]); err == nil {
		product := ps.GetOne(pid)
		if product == nil {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "error encoding data", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "invalid data in the url", http.StatusBadRequest)

}
