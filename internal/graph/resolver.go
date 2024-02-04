package graph

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/model/lov_field_model"
	"go-gin-example/internal/model/lov_model"
	"go-gin-example/internal/model/sys_user_model"
	"go-gin-example/internal/service/lov"
	"go-gin-example/internal/service/lov_field"
	"go-gin-example/internal/service/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService     user.Service
	LovService      lov.Service
	LovFieldService lov_field.Service
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func OfUser(sysUser *sys_user_model.SysUser) *graph_model.User {
	return &graph_model.User{
		ID:        int(sysUser.ID),
		Username:  sysUser.Username,
		CreatedAt: int(sysUser.CreatedAt.Unix()),
		Nickname:  &sysUser.Nickname,
		Phone:     &sysUser.Phone,
		Gender:    int(sysUser.Gender),
		Head:      &sysUser.Head,
		Remark:    &sysUser.Remark,
		State:     int(sysUser.State),
		Email:     &sysUser.Email,
	}
}

func OfUsers(sysUsers *[]sys_user_model.SysUser) []*graph_model.User {
	var users []*graph_model.User
	for _, sysUser := range *sysUsers {
		users = append(users, OfUser(&sysUser))
	}
	return users
}

func OfLov(lov *lov_model.Lov) *graph_model.Lov {
	return &graph_model.Lov{
		ID:   int(lov.ID),
		Code: lov.Code,
		Name: lov.Name,
		Desc: &lov.Desc,
	}
}

func OfLovList(lovList *[]lov_model.Lov) []*graph_model.Lov {
	var list []*graph_model.Lov
	for _, lov := range *lovList {
		list = append(list, OfLov(&lov))
	}
	return list
}

func OfLovField(lovField *lov_field_model.LovField) *graph_model.LovField {
	return &graph_model.LovField{
		ID:     int(lovField.ID),
		Label:  lovField.Label,
		Value:  int(lovField.Value),
		Desc:   &lovField.Desc,
		Status: int(lovField.Stauts),
	}
}

func OfLovFields(lovFields *[]lov_field_model.LovField) []*graph_model.LovField {
	var fields []*graph_model.LovField
	for _, lovField := range *lovFields {
		fields = append(fields, OfLovField(&lovField))
	}
	return fields
}
