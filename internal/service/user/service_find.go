package user

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/sys_user_model"
)

func (s *service) FindById(userId int64) (*sys_user_model.SysUser, error) {
	return sys_user_model.FindById(userId)
}

func (s *service) FindList(pagination *graph_model.Pagination, input *graph_model.UsersInput) (*[]sys_user_model.SysUser, error) {
	return sys_user_model.FindList(pagination, input)
}

func (s *service) FindTotal(input *graph_model.UsersInput) (int64, error) {
	return sys_user_model.FindTotal(input)
}
