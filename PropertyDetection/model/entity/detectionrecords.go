package entity

import "time"

type DetectionRecords struct {
	Id
	Status
	Remark
	Deleted
	Url        string    `json:"url" gorm:"column:url" column:"url" comment:"url"`
	Score      string    `json:"score" gorm:"column:score" column:"score" comment:"score"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time" column:"create_time" comment:"create_time"`
}

func (DetectionRecords) TableName() string {
	return "detection_records"
}
