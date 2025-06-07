package entity

type Role struct {
	Id
	Status
	Remark
	Deleted
	Name string `json:"name" gorm:"column:name" column:"name" comment:"角色名"`
}

func (Role) TableName() string {
	return "role"
}
