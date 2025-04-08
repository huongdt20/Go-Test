package storage

import (
	"Go-Test/pkg/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *Storage) GetCategoryByID(ctx context.Context, id int64) (*model.Category, error) {
	category := &model.Category{}
	if err := s.DB.WithContext(ctx).Where("id = ?", id).First(category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errorDataNotFound)
		}
		return nil, err
	}
	return category, nil
}
