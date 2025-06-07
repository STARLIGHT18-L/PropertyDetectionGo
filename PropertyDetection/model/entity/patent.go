package entity

import "time"

type Patent struct {
	Id
	Status
	Remark
	Deleted
	Url          string    `json:"url" gorm:"column:url" column:"url" comment:"专利地址"`
	Name         string    `json:"name" gorm:"column:name" column:"name" comment:"专利名称"`
	Content      string    `json:"content" gorm:"column:content" column:"content" comment:"专利内容"`
	Vector       string    `json:"vector" gorm:"column:vector" column:"vector" comment:"专利向量"`
	Owner        string    `json:"owner" gorm:"column:owner" column:"owner" comment:"专利拥有者"`
	RegisterDate time.Time `json:"registerDate" gorm:"column:register_date" column:"register_date" comment:"专利注册时间"`
}

func (Patent) TableName() string {
	return "patent"
}
