package entity

import (
	"strconv"
)

type User struct {
	Id
	Status
	Remark
	Deleted
	Username string `json:"username" gorm:"column:username" column:"username" comment:"用户名"`
	Password string `json:"password" gorm:"column:password" column:"password" comment:"密码"`
	Email    string `json:"email" gorm:"column:email" column:"email" comment:"邮箱"`
	RoleId   int    `json:"roleId" gorm:"column:role_id" column:"role_id" comment:"角色ID"`
}

func (User) TableName() string {
	return "user"
}

// token
func (user User) GetUid() string {
	return strconv.Itoa(user.Id.Id)
}
