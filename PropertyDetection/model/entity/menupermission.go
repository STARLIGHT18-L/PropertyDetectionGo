package entity

type MenuPermission struct {
	Id
	Status
	Remark
	Deleted
	MenuId int `json:"menuId" gorm:"column:menu_id" column:"menu_id" comment:"菜单ID"`
	RoleId int `json:"roleId" gorm:"column:role_id" column:"role_id" comment:"角色ID"`
}

func (MenuPermission) TableName() string {
	return "menu_permission"
}
