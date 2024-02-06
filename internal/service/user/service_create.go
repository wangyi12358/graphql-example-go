package user

import (
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/sys_user_model"
	"go-gin-example/pkg/util/passwrod"
	"go-gin-example/pkg/util/string_util"
)

func (s *service) Create(input *graph_model.CreateUser) (*sys_user_model.SysUser, error) {
	user := &sys_user_model.SysUser{
		Username: input.Username,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		Gender:   int16(input.Gender),
		Remark:   string_util.GetStringFromPointer(input.Remark, ""),
		Email:    input.Email,
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
