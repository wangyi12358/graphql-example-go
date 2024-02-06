package sys_user_model

import (
	"fmt"
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/pkg/model"
	"gorm.io/gorm"
)

func (s *SysUser) MakeSalt() {
	s.Salt = "123"
}

func Login(username string, password string) error {
	var count int64
	err := model.DB.Model(&SysUser{}).Where(&SysUser{
		Username: username,
		Password: password,
	}).Count(&count).Error
	if err != nil {
		fmt.Println("1")
		return err
	}
	if count == 0 {
		return fmt.Errorf("用户名或密码错误")
	}
	return nil
}

func FindByUsername(username string) (*SysUser, error) {
	var user SysUser
	if err := model.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindById(id int64) (*SysUser, error) {
	var user SysUser
	if err := model.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func connectUserQuery(input *graph_model.UsersInput) *gorm.DB {
	db := model.DB
	if input == nil {
		return db
	}
	if input.Nickname != nil {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", *input.Nickname))
	}
	return db
}

func FindList(pagination *graph_model.Pagination, input *graph_model.UsersInput) (*[]SysUser, error) {
	var users []SysUser
	offset := (pagination.PageSize - 1) * pagination.Current
	if err := connectUserQuery(input).
		Find(&users).
		Offset(offset).
		Limit(pagination.PageSize).
		Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func FindTotal(input *graph_model.UsersInput) (int64, error) {
	var total int64
	if err := connectUserQuery(input).
		Model(&SysUser{}).
		Count(&total).
		Error; err != nil {
		return 0, err
	}
	return total, nil
}

func Create(user *SysUser) error {
	if err := model.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
