package user

import (
	"fmt"
	"go-gin-example/internal/model/sys_user_model"
	token2 "go-gin-example/pkg/token"
	"go-gin-example/pkg/util/passwrod"
)

func (s *service) Login(username string, passwordStr string) (token *string, user *sys_user_model.SysUser, err error) {
	user, err = sys_user_model.FindByUsername(username)
	if err != nil {
		return
	}
	fmt.Printf("user: %v\n", user)
	hashedPassword, err := passwrod.GeneratePassword(passwordStr, user)
	if err != nil {
		return
	}
	fmt.Printf("hashedPassword: %s\n", hashedPassword)
	err = sys_user_model.Login(username, hashedPassword)
	if err != nil {
		return
	}
	token, err = token2.GenerateToken(user.ID)
	if err != nil {
		return
	}
	return
}
