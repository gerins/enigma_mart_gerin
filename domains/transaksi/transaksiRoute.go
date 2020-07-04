package transaction

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitTransactionRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	TransactionController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.HandleFunc("", TransactionController.HandleGETAllTransactions()).Methods("GET")
	p.HandleFunc("/montly/{month}", TransactionController.HandleGETAllTransactionsMontly()).Methods("GET")
	p.HandleFunc("/daily/{daily}", TransactionController.GetTransactionsDaily()).Methods("GET")
	p.HandleFunc("/{id}", TransactionController.HandleGETTransaction()).Methods("GET")
	p.HandleFunc("", TransactionController.HandlePOSTTransactions()).Methods("POST")
	p.HandleFunc("/{id}", TransactionController.HandleUPDATETransactions()).Methods("PUT")
	p.HandleFunc("/{id}", TransactionController.HandleDELETETransactions()).Methods("DELETE")
}
