package lov

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/lov_model"
	"gorm.io/gorm"
)

type Service interface {
	List(pagination *graph_model.Pagination, input *graph_model.LovPageInput) (*[]lov_model.Lov, error)
	Total(input *graph_model.LovPageInput) (int64, error)
	Create(input *graph_model.CreateLov) (*lov_model.Lov, error)
}

type service struct {
	db *gorm.DB
}

func (s *service) List(pagination *graph_model.Pagination, input *graph_model.LovPageInput) (*[]lov_model.Lov, error) {
	list, err := lov_model.List(pagination, input)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) Total(input *graph_model.LovPageInput) (int64, error) {
	total, err := lov_model.Total(input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *service) Create(input *graph_model.CreateLov) (*lov_model.Lov, error) {
	lov := &lov_model.Lov{
		Name: input.Name,
		Code: input.Code,
		Desc: *input.Desc,
	}
	err := lov_model.Create(lov)
	if err != nil {
		return nil, err
	}
	return lov, nil
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
