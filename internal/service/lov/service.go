package lov

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/lov_field_model"
	"go-gin-example/internal/model/lov_model"
	"go-gin-example/pkg/util/string_util"
	"gorm.io/gorm"
)

type Service interface {
	FindById(id int64) (*lov_model.Lov, error)
	List(pagination *graph_model.Pagination, input *graph_model.LovPageInput) (*[]lov_model.Lov, error)
	Total(input *graph_model.LovPageInput) (int64, error)
	Create(input *graph_model.CreateLov) (*lov_model.Lov, error)
	DeleteById(id int64) error
}

type service struct {
	db *gorm.DB
}

func (s *service) FindById(id int64) (*lov_model.Lov, error) {
	var lov *lov_model.Lov
	err := s.db.First(&lov, id).Error
	if err != nil {
		return nil, err
	}
	return lov, nil
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
		Desc: string_util.GetStringFromPointer(input.Desc, ""),
	}
	err := lov_model.Create(lov)
	if err != nil {
		return nil, err
	}
	return lov, nil
}

func (s *service) DeleteById(id int64) error {
	db := s.db.Begin()
	err := db.Delete(&lov_model.Lov{}, id).Error
	if err != nil {
		db.Rollback()
		return err
	}
	err = db.Where("lov_id = ?", id).Delete(&lov_field_model.LovField{}).Error
	if err != nil {
		db.Rollback()
		return err
	}
	return db.Commit().Error
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
