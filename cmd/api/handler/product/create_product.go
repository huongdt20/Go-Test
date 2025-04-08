package product

import (
	"Go-Test/cmd/api/dto"
	"Go-Test/pkg/model"
	"Go-Test/pkg/storage"
	util2 "Go-Test/pkg/util"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type createProduct struct {
	ds *storage.Storage
}

func NewCreateProductHandler(ds *storage.Storage) http.Handler {
	return &createProduct{
		ds: ds,
	}
}

func (h *createProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := h.parseRequest(r)
	if err != nil {
		util2.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	ctx := r.Context()
	err = h.validateRequest(ctx, req)
	if err != nil {
		util2.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.ds.Transaction(func(subds *storage.Storage) error {
		product, err := subds.CreateProduct(ctx, &model.Product{
			Name:          req.ProductName,
			Description:   req.Description,
			Status:        model.StatusInStock,
			Price:         req.Price,
			StockQuantity: req.Quantity,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})
		if err != nil {
			return err
		}

		err = subds.CreateProductCategory(ctx, &model.ProductCategory{
			ProductID:  product.ID,
			CategoryID: req.CategoryID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		util2.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	util2.ResponseSuccess(w, nil)
}

func (h *createProduct) parseRequest(r *http.Request) (*dto.CreateProductRequest, error) {
	req := &dto.CreateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (h *createProduct) validateRequest(ctx context.Context, req *dto.CreateProductRequest) error {
	if req.ProductName == "" {
		return errors.New("product name is a required field")
	}

	if req.Price <= 0 {
		return errors.New("product price format is invalid")
	}

	if req.Quantity <= 0 {
		return errors.New("quantity invalid")
	}

	if req.CategoryID <= 0 {
		return errors.New("category invalid")
	}

	_, err := h.ds.GetCategoryByID(ctx, req.CategoryID)
	if err != nil {
		return errors.New("category invalid")
	}
	return nil
}
