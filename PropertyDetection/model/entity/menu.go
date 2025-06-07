package entity

type Menu struct {
	Id
	Status
	Remark
	Deleted
	Label       string `json:"label" gorm:"column:label" comment:"标签"`
	Path        string `json:"path" gorm:"column:path" comment:"路径"`
	Component   string `json:"component" gorm:"column:component" comment:"组件"`
	Icon        string `json:"icon" gorm:"column:icon" comment:"图标，默认值为 icon-caidan"`
	IconBgColor string `json:"iconBgColor" gorm:"column:icon_bg_color" comment:"图标背景颜色，默认值为 #fff"`
	Meta        string `json:"meta" gorm:"column:meta" comment:"元数据"`
	Href        string `json:"href" gorm:"column:href" comment:"链接"`
	ParentId    int    `json:"parentId" gorm:"column:parent_id" comment:"父菜单 ID"`
	Children    []Menu `json:"children" gorm:"-" comment:"子菜单列表"`
}

func (Menu) TableName() string {
	return "menu"
}
