package route

import (
	"Go-Test/cmd/api/handler/product"
	"Go-Test/pkg/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(ds *storage.Storage) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/products", product.NewCreateProductHandler(ds).ServeHTTP).Methods(http.MethodPost)
	//r.HandleFunc("/products", product.NewGetProductsHandler(ds).ServeHTTP).Methods(http.MethodGet)

	return r
}
