package entity

type Image struct {
	Id
	Status
	Remark
	Deleted
	Url            string `json:"url" gorm:"column:url" column:"url" comment:"图片地址"`
	Name           string `json:"name" gorm:"column:name" column:"name" comment:"图片名称"`
	Content        string `json:"content" gorm:"column:content" column:"content" comment:"描述"`
	Vector         string `json:"vector" gorm:"column:vector" column:"vector" comment:"图片向量"`
	RelationPatent string `json:"relationPatent" gorm:"column:relation_patent" column:"relation_patent" comment:"图片关系"`
}

func (Image) TableName() string {
	return "image"
}
