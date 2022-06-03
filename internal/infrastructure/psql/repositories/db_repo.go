package repositories

import (
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(dbConn *gorm.DB) *DBRepository {
	return &DBRepository{
		db: dbConn,
	}
}

func (r *DBRepository) DB() *gorm.DB {
	return r.db
}
