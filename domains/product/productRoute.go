package product

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitProductRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	productController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", productController.HandleGETAllProducts()).Methods("GET")
	p.HandleFunc("/{id}", productController.HandleGETProduct()).Methods("GET")
	p.HandleFunc("", productController.HandlePOSTProducts()).Methods("POST")
	p.HandleFunc("/{id}", productController.HandleUPDATEProducts()).Methods("PUT")
	p.HandleFunc("/{id}", productController.HandleDELETEProducts()).Methods("DELETE")
}
