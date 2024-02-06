package user

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/sys_user_model"
	"gorm.io/gorm"
)

type Service interface {
	Create(input *graph_model.CreateUser) (*sys_user_model.SysUser, error)
	Login(username string, password string) (token *string, user *sys_user_model.SysUser, err error)
	FindById(userId int64) (*sys_user_model.SysUser, error)
	FindList(pagination *graph_model.Pagination, input *graph_model.UsersInput) (*[]sys_user_model.SysUser, error)
	FindTotal(input *graph_model.UsersInput) (int64, error)
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
