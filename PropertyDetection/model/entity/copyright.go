package entity

import "time"

type Copyright struct {
	Id
	Status
	Remark
	Deleted
	Url          string    `json:"url" gorm:"column:url" column:"url" comment:"图片地址"`
	Name         string    `json:"name" gorm:"column:name" column:"name" comment:"图片名称"`
	Content      string    `json:"content" gorm:"column:content" column:"content" comment:"描述"`
	Vector       string    `json:"vector" gorm:"column:vector" column:"vector" comment:"图片向量"`
	Type         string    `json:"type" gorm:"column:type" column:"type" comment:"图片类型"`
	Owner        string    `json:"owner" gorm:"column:owner" column:"owner" comment:"图片拥有者"`
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date" column:"register_date" comment:"图片注册时间"`
}

func (Copyright) TableName() string {
	return "copyright"
}
