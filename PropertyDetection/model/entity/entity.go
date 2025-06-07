package entity

import (
	"time"
)

type Entity struct {
	Id      int    `json:"id" gorm:"column:id" column:"id" comment:"ID主键"`
	Deleted int    `json:"deleted" gorm:"column:deleted" column:"deleted" comment:"逻辑删除"`
	Status  int    `json:"status" gorm:"column:status" column:"status" comment:"状态"`
	Remark  string `json:"remark" gorm:"column:remark" column:"remark" comment:"备注"`
}
type Id struct {
	Id int `json:"id" gorm:"column:id" column:"id" comment:"ID主键"`
}
type Status struct {
	Status int `json:"status" gorm:"column:status" column:"status" comment:"状态"`
}
type Deleted struct {
	Deleted int `json:"deleted" gorm:"column:deleted" column:"deleted" comment:"逻辑删除"`
}
type Remark struct {
	Remark string `json:"remark" gorm:"column:remark" column:"remark" comment:"备注"`
}

func PrimaryKey() int {
	return int(time.Now().Unix())
}
