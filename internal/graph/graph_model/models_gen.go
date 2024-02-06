// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph_model

type CreateLov struct {
	// 名称
	Name string `json:"name"`
	// 编码
	Code string `json:"code"`
	// 描述
	Desc *string `json:"desc,omitempty"`
}

type CreateLovField struct {
	// 名称
	Label string `json:"label"`
	// 描述
	Desc *string `json:"desc,omitempty"`
	// 值
	Value int `json:"value"`
	// 状态
	Status int `json:"status"`
	// lovId
	LovID int `json:"lovId"`
}

type CreateUser struct {
	Username string  `json:"username"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password" validate:"max=16,min=6" name:"密码"`
	Phone    string  `json:"phone"`
	Gender   int     `json:"gender"`
	Remark   *string `json:"remark,omitempty"`
	Email    string  `json:"email" validate:"email" name:"邮箱"`
}

type Login struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Lov struct {
	ID int `json:"id"`
	// 名称
	Name string `json:"name"`
	// 编码
	Code string `json:"code"`
	// 描述
	Desc *string `json:"desc,omitempty"`
}

type LovField struct {
	ID int `json:"id"`
	// 名称
	Label string `json:"label"`
	// 描述
	Desc *string `json:"desc,omitempty"`
	// 值
	Value int `json:"value"`
	// 状态
	Status int `json:"status"`
}

type LovPage struct {
	// 总数
	Total int `json:"total"`
	// 列表
	Data []*Lov `json:"data"`
}

type LovPageInput struct {
	// 名称模糊查询
	Name *string `json:"name,omitempty"`
	// 编码模糊查询
	Code *string `json:"code,omitempty"`
}

type Mutation struct {
}

type Pagination struct {
	// 页码
	Current int `json:"current"`
	// 每页数量
	PageSize int `json:"pageSize"`
}

type Query struct {
}

type User struct {
	ID int `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 创建时间
	CreatedAt int `json:"createdAt"`
	// 昵称
	Nickname *string `json:"nickname,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 性别
	Gender int `json:"gender"`
	// 头像
	Head *string `json:"head,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 状态
	State int `json:"state"`
	// 邮箱
	Email *string `json:"email,omitempty"`
}

type Users struct {
	// 总数
	Total int `json:"total"`
	// 列表
	Data []*User `json:"data"`
}

type UsersInput struct {
	Nickname *string `json:"nickname,omitempty"`
}
