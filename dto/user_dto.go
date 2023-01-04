package dto

import "junsonjack.cn/go_vue/model"

// DTO:数据传输对象
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto{
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}
