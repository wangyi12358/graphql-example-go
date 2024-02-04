package lov_field

import (
	"gorm.io/gorm"
)

type Service interface {
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
