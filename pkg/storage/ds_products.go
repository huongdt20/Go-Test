package storage

import (
	"Go-Test/pkg/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *Storage) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := s.DB.WithContext(ctx).Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, err
}

func (s *Storage) GetProductByProductName(ctx context.Context, productName string) (*model.Product, error) {
	product := &model.Product{}
	err := s.DB.WithContext(ctx).Model(&model.Product{}).Where("name = ?", productName).First(product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errorDataNotFound)
		}
		return nil, err
	}
	return product, nil
}
