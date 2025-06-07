package mapper

import (
	"PropertyDetection/config"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"fmt"
)

// UserMapper 结构体用于封装用户相关的数据库操作
type UserMapper struct{}

// GetUserByUsername 根据用户名查询用户信息
func (mapper UserMapper) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := config.Boot.Mysql.Where("username = ? AND deleted=0", username).First(&user).Error
	return user, err
}

// GetRoleNamesByRoleId 根据角色 ID 查询角色名称列表
func (mapper UserMapper) GetRoleNamesByRoleId(roleId int) ([]string, error) {
	var roles []entity.Role
	var roleNames []string
	err := config.Boot.Mysql.Where("id = ? AND deleted = 0", roleId).Select("name").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	return roleNames, nil
}

// GetMenuListByRoleId 根据角色 ID 获取菜单列表
func (mapper UserMapper) GetMenuListByRoleId(roleId int) ([]entity.Menu, error) {
	// 定义 SQL 查询语句，使用 CTE 进行递归查询
	sql := `
    WITH RECURSIVE menu_tree AS (
        SELECT id, deleted, status, remark, label, path, component, icon, icon_bg_color, meta, href, parent_id
        FROM menu
        WHERE parent_id = 0 AND deleted = 0 AND status = 1
        UNION ALL
        SELECT m.id, m.deleted, m.status, m.remark, m.label, m.path, m.component, m.icon, m.icon_bg_color, m.meta, m.href, m.parent_id
        FROM menu m
        JOIN menu_tree mt ON m.parent_id = mt.id
        WHERE m.deleted = 0 AND m.status = 1
    )
    SELECT DISTINCT t.id, t.deleted, t.status, t.remark, t.label, t.path, t.component, t.icon, t.icon_bg_color, t.meta, t.href, t.parent_id
    FROM menu_tree t
    LEFT JOIN menu_permission mp ON t.id = mp.menu_id
    WHERE mp.role_id = ?
    `
	var menuList []entity.Menu
	err := config.Boot.Mysql.Raw(sql, roleId).Scan(&menuList).Error // 执行 SQL 查询
	if err != nil {
		return nil, err
	}
	// 构建关系映射
	relationMap := make(map[int][]int)
	menuMap := make(map[int]entity.Menu)
	for _, menu := range menuList {
		menuMap[menu.Id.Id] = menu
		relationMap[menu.ParentId] = append(relationMap[menu.ParentId], menu.Id.Id)
	}
	// 构建菜单树
	return buildMenus(relationMap[0], relationMap, menuMap), nil
}
func buildMenus(ids []int, relations map[int][]int, menuMap map[int]entity.Menu) []entity.Menu {
	menus := make([]entity.Menu, len(ids))
	for i, id := range ids {
		menu := menuMap[id]
		if childIDs, ok := relations[id]; ok {
			menu.Children = buildMenus(childIDs, relations, menuMap)
		}
		menus[i] = menu
	}
	return menus
}
func (mapper UserMapper) Insert(user *entity.User) error {
	return config.Boot.Mysql.Create(user).Error
}
func (mapper UserMapper) Update(user *entity.User) error {
	return config.Boot.Mysql.Model(user).Updates(user).Error
}
func (mapper UserMapper) Delete(user *entity.User) error {
	return config.Boot.Mysql.Model(user).Where("id = ?", user.Id.Id).Update("deleted", 1).Error
}
func (mapper UserMapper) Page(page *base.Page, user *entity.User, info *entity.User) error {
	// 构建查询条件
	query := config.Boot.Mysql.Select("id, username, email, role_id, status, remark, deleted").Where("deleted = ?", 0)
	if user.Username != "" {
		query = query.Where("username LIKE ?", fmt.Sprintf("%%%s%%", user.Username))
	}
	if user.Email != "" {
		query = query.Where("email LIKE ?", fmt.Sprintf("%%%s%%", user.Email))
	}
	if user.Remark.Remark != "" {
		query = query.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", user.Remark.Remark))
	}
	if info.RoleId > 0 {
		query = query.Where("role_id >= ?", info.RoleId)
	}
	if user.RoleId > 0 {
		query = query.Where("role_id = ?", user.RoleId)
	}
	if user.Status.Status > 0 {
		query = query.Where("status = ?", user.Status.Status)
	}
	var total int64 // 计算总数
	if err := query.Model(user).Count(&total).Error; err != nil {
		return err
	}
	var users []entity.User // 分页查询
	if err := query.Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&users).Error; err != nil {
		return err
	}
	// 填充分页信息
	page.Total = total
	page.Records = users
	return nil
}
