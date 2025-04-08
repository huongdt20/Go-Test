package storage

import (
	cutil "Go-Test/pkg/util"
	"go.uber.org/multierr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const errorDataNotFound = "data not found"

type Storage struct {
	*gorm.DB
	cache Cache
}

func NewStorage() (*Storage, error) {
	cfg := cutil.LoadConfig()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName:           "pgx",
		DSN:                  cfg.DSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error connect: %v", err)
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return &Storage{
		DB:    db,
		cache: &CacheMock{},
	}, nil
}

func (s *Storage) Close() {
	if s.DB != nil {
		sqlDB, _ := s.DB.DB()
		sqlDB.Close()
	}
}

func (s *Storage) Transaction(f func(subds *Storage) error) (err error) {
	transDB := s.DB.Begin()
	if err = f(&Storage{
		DB:    transDB,
		cache: s.cache,
	}); err != nil {
		if rErr := transDB.Rollback().Error; rErr != nil {
			return multierr.Combine(err, rErr)
		}
		return err
	}

	err = transDB.Commit().Error
	return err
}
