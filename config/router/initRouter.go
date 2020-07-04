package router

import (
	"database/sql"
	"enigma_mart_gerin/domains/category"
	"enigma_mart_gerin/domains/product"
	transaction "enigma_mart_gerin/domains/transaksi"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PRODUCT_MAIN_ROUTE     = "/products"
	CATEGORY_MAIN_ROUTE    = "/category"
	TRANSACTION_MAIN_ROUTE = "/transaction"
)

type ConfigRouter struct {
	DB     *sql.DB
	Router *mux.Router
}

func (ar *ConfigRouter) InitRouter() {
	product.InitProductRoute(PRODUCT_MAIN_ROUTE, ar.DB, ar.Router)
	category.InitCategoryRoute(CATEGORY_MAIN_ROUTE, ar.DB, ar.Router)
	transaction.InitTransactionRoute(TRANSACTION_MAIN_ROUTE, ar.DB, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)

}

// NewAppRouter for creating new Route
func NewAppRouter(db *sql.DB, r *mux.Router) *ConfigRouter {
	return &ConfigRouter{
		DB:     db,
		Router: r,
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}
