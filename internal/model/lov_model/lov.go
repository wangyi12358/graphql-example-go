package lov_model

import (
	"fmt"
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/pkg/model"
	"gorm.io/gorm"
)

func connectLovQuery(input *graph_model.LovPageInput) *gorm.DB {
	db := model.DB
	if input == nil {
		return db
	}
	if input.Name != nil {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *input.Name))
	}
	if input.Code != nil {
		db = db.Where("code LIKE ?", fmt.Sprintf("%%%s%%", *input.Code))
	}
	return db
}

func List(pagination *graph_model.Pagination, input *graph_model.LovPageInput) (*[]Lov, error) {
	var lovList []Lov
	offset := (pagination.PageSize - 1) * pagination.Current
	if err := connectLovQuery(input).
		Find(&lovList).
		Offset(offset).
		Limit(pagination.PageSize).
		Error; err != nil {
		return nil, err
	}
	return &lovList, nil
}

func Total(input *graph_model.LovPageInput) (int64, error) {
	var total int64
	if err := connectLovQuery(input).Model(&Lov{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func Create(lov *Lov) error {
	return model.DB.Create(lov).Error
}

func DeleteById(id int64) error {
	return model.DB.Delete(&Lov{}, id).Error
}
