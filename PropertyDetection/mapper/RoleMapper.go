package mapper

import (
	"PropertyDetection/config"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"fmt"
	"strconv"
)

type RoleMapper struct {
}

func (mapper RoleMapper) Insert(role *entity.Role) error {
	return config.Boot.Mysql.Create(role).Error
}
func (mapper RoleMapper) Update(role *entity.Role) error {
	return config.Boot.Mysql.Model(role).Updates(role).Error
}
func (mapper RoleMapper) Delete(role *entity.Role) error {
	return config.Boot.Mysql.Model(role).Update("deleted", 1).Error
}
func (mapper RoleMapper) Page(page *base.Page, role *entity.Role, info *entity.User) error {
	query := config.Boot.Mysql.Select("id, name, remark, deleted, status").Where("deleted = ?", 0)
	if role.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", role.Name))
	}
	if role.Remark.Remark != "" {
		query = query.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", role.Remark.Remark))
	}
	if info.RoleId > 0 {
		query = query.Where("id >= ?", info.RoleId)
	}
	var total int64 // 计算总数
	if err := query.Model(role).Count(&total).Error; err != nil {
		return err
	}
	var roles []entity.Role // 分页查询
	if err := query.Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&roles).Error; err != nil {
		return err
	}
	// 填充分页信息
	page.Total = total
	page.Records = roles
	return nil
}
func (mapper RoleMapper) ListMaps(user *entity.User) []entity.Role {
	roleList := []entity.Role{}
	config.Boot.Mysql.Select("id, name").Where("deleted = ? AND id >= ?", 0, user.RoleId).Find(&roleList)
	return roleList
}
func (mapper RoleMapper) SetMenuPermission(roleId int, roleName string, ids []interface{}) error {
	// 查询已经添加过的菜单
	var menuPermissions []entity.MenuPermission
	config.Boot.Mysql.Select("id, menu_id, role_id").Where("role_id =? AND deleted =?", roleId, 0).Find(&menuPermissions)
	menuPermissionMap := make(map[int]bool)
	for _, mp := range menuPermissions {
		menuPermissionMap[mp.MenuId] = true
	}
	tx := config.Boot.Mysql.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	for _, o := range ids {
		id, err := strconv.Atoi(fmt.Sprintf("%v", o))
		if err != nil {
			tx.Rollback()
			return err
		}
		if !menuPermissionMap[id] {
			mp := entity.MenuPermission{
				Id:      entity.Id{Id: entity.PrimaryKey()},
				Status:  entity.Status{Status: 1},
				Deleted: entity.Deleted{Deleted: 0},
				Remark:  entity.Remark{Remark: roleName},
				MenuId:  id,
				RoleId:  roleId,
			}
			if err = tx.Create(&mp).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
