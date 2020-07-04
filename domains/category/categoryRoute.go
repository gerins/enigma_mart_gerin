package category

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitCategoryRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	CategoryController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", CategoryController.HandleGETAllCategorys()).Methods("GET")
	p.HandleFunc("/{id}", CategoryController.HandleGETCategory()).Methods("GET")
	p.HandleFunc("", CategoryController.HandlePOSTCategorys()).Methods("POST")
	p.HandleFunc("/{id}", CategoryController.HandleUPDATECategorys()).Methods("PUT")
	p.HandleFunc("/{id}", CategoryController.HandleDELETECategorys()).Methods("DELETE")
}
