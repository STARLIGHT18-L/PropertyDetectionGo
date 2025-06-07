package entity

import "time"

type Trademark struct {
	Id
	Status
	Remark
	Deleted
	Url          string    `json:"url" gorm:"column:url" column:"url" comment:"商标地址"`
	Name         string    `json:"name" gorm:"column:name" column:"name" comment:"商标名称"`
	Content      string    `json:"content" gorm:"column:content" column:"content" comment:"商标内容"`
	Vector       string    `json:"vector" gorm:"column:vector" column:"vector" comment:"商标向量"`
	Owner        string    `json:"owner" gorm:"column:owner" column:"owner" comment:"商标拥有者"`
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date" column:"register_date" comment:"商标注册时间"`
}

func (Trademark) TableName() string {
	return "trademark"
}
