package product

import (
	"Go-Test/pkg/storage"
	"net/http"
)

type getProductsHandler struct {
	ds *storage.Storage
}

func NewGetProductsHandler(s *storage.Storage) http.Handler {
	return &getProductsHandler{
		ds: s,
	}
}

func (h *getProductsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
