package user

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/sys_user_model"
	"go-gin-example/pkg/util/copy_struct"
	"go-gin-example/pkg/util/passwrod"
)

func (s *service) Create(input *graph_model.CreateUser) (*sys_user_model.SysUser, error) {
	user := &sys_user_model.SysUser{}
	err := copy_struct.CopyStruct(input, user)
	if err != nil {
		return nil, err
	}
	salt, err := passwrod.GenerateRandomSalt()
	if err != nil {
		return nil, err
	}
	user.Salt = salt
	hashedPassword, err := passwrod.GeneratePassword(input.Password, user)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	err = sys_user_model.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
