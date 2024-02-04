// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package product_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameProduct = "product"

// Product mapped from table <product>
type Product struct {
	ID         int64          `gorm:"column:id;primaryKey" json:"id"`
	Name       string         `gorm:"column:name" json:"name"`
	URL        string         `gorm:"column:url" json:"url"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CategoryID int64          `gorm:"column:category_id" json:"category_id"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}