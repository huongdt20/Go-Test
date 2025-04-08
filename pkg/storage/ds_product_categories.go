package storage

import (
	"Go-Test/pkg/model"
	"context"
)

func (s *Storage) CreateProductCategory(ctx context.Context, productCategory *model.ProductCategory) error {
	return s.DB.WithContext(ctx).Create(productCategory).Error
}
