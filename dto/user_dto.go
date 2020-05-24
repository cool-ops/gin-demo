package dto

import "github.com/cool-ops/gin-demo/model"

// 定义用户传输对象结构体
type UserDTO struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDTO(user model.User) UserDTO{
	return UserDTO{
		Name:      user.UserName,
		Telephone: user.Telephone,
	}
}