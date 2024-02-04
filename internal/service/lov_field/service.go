package lov_field

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/lov_field_model"
	"go-gin-example/pkg/util/string_util"
	"gorm.io/gorm"
)

type Service interface {
	FindListByLovId(lovId int64) (*[]lov_field_model.LovField, error)
	CreateLovField(input *graph_model.CreateLovField) (*lov_field_model.LovField, error)
}

type service struct {
	db *gorm.DB
}

func (s *service) FindListByLovId(lovId int64) (*[]lov_field_model.LovField, error) {
	var list []lov_field_model.LovField
	err := s.db.Where("lov_id = ?", lovId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *service) CreateLovField(input *graph_model.CreateLovField) (*lov_field_model.LovField, error) {
	lovField := &lov_field_model.LovField{
		LovID:  int64(input.LovID),
		Label:  input.Label,
		Value:  int16(input.Value),
		Stauts: 1,
		Desc:   string_util.GetStringFromPointer(input.Desc, ""),
	}
	err := s.db.Create(lovField).Error
	if err != nil {
		return nil, err
	}
	return lovField, nil
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
