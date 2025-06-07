package mapper

import (
	"PropertyDetection/config"
	"PropertyDetection/model/entity"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type MenuMapper struct {
}

// GetMenuList 获取菜单列表
func (mapper MenuMapper) GetMenuList(m *entity.Menu) ([]entity.Menu, error) {
	// 构建 CTE 查询语句
	query := `
	WITH RECURSIVE menu_tree AS (
		SELECT id, deleted, status, remark, label, path, component, icon, icon_bg_color, meta, href, parent_id
		FROM menu
		WHERE deleted = 0 AND status = 1
	`
	if m.Id.Id != 0 {
		query += fmt.Sprintf(" AND id = %d", m.Id.Id)
	} else {
		query += " AND parent_id = 0"
	}
	query += `
		UNION ALL
		SELECT m.id, m.deleted, m.status, m.remark, m.label, m.path, m.component, m.icon, m.icon_bg_color, m.meta, m.href, m.parent_id
		FROM menu m
		JOIN menu_tree mt ON m.parent_id = mt.id
		WHERE m.deleted = 0 AND m.status = 1
	)
	SELECT DISTINCT id, deleted, status, remark, label, path, component, icon, icon_bg_color, meta, href, parent_id
	FROM menu_tree
	`
	if m.Remark.Remark != "" {
		query += fmt.Sprintf(" WHERE remark LIKE '%%%s%%'", m.Remark.Remark)
	}
	var menuList []entity.Menu
	if err := config.Boot.Mysql.Raw(query).Scan(&menuList).Error; err != nil {
		return nil, err
	}
	// 构建菜单关系映射
	relations := make(map[int][]entity.Menu)
	for _, menu := range menuList {
		relations[menu.ParentId] = append(relations[menu.ParentId], menu)
	}
	// 构建菜单树
	result := buildMenuTree(relations[0], relations)
	// 获取模糊查询条件
	labelFilter := m.Label
	if labelFilter == "" {
		return result, nil
	}
	// 对名称进行过滤
	result = filterLabel(result, labelFilter)
	return result, nil
}

// buildMenuTree 递归构建菜单树
func buildMenuTree(menus []entity.Menu, relations map[int][]entity.Menu) []entity.Menu {
	for i := range menus {
		if childMenus, ok := relations[menus[i].Id.Id]; ok {
			menus[i].Children = buildMenuTree(childMenus, relations)
		}
	}
	return menus
}

// filterLabel 递归过滤菜单列表
func filterLabel(menus []entity.Menu, labelFilter string) []entity.Menu {
	result := make([]entity.Menu, 0)
	for _, menu := range menus {
		isMatch := strings.Contains(menu.Label, labelFilter)
		filteredChildren := filterLabel(menu.Children, labelFilter)
		if isMatch {
			// 如果当前名称匹配，包含所有子节点，无论是否匹配
			menu.Children = filteredChildren
			result = append(result, menu)
		} else if len(filteredChildren) > 0 {
			// 如果当前名称不匹配，但有匹配的子节点，直接将匹配的子节点添加到结果中
			menu.Children = filteredChildren
			result = append(result, menu)
		}
	}
	return result
}
func (mapper MenuMapper) Insert(menu *entity.Menu, mp *entity.MenuPermission) error {
	return config.Boot.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(menu).Error; err != nil {
			return err
		}
		return tx.Create(mp).Error
	})
}
func (mapper MenuMapper) Update(menu *entity.Menu) error {
	return config.Boot.Mysql.Model(menu).Updates(menu).Error
}
func (mapper MenuMapper) Delete(menu *entity.Menu) error {
	// 构建 CTE 查询语句
	query := `
    WITH RECURSIVE menu_tree AS (
        SELECT id, deleted, status, remark, label, path, component, icon, icon_bg_color, meta, href, parent_id
        FROM menu
        WHERE deleted = 0 AND status = 1
    `
	if menu.Id.Id != 0 {
		query += fmt.Sprintf(" AND id = %d", menu.Id.Id)
	} else {
		query += " AND parent_id = 0"
	}
	query += `
        UNION ALL
        SELECT m.id, m.deleted, m.status, m.remark, m.label, m.path, m.component, m.icon, m.icon_bg_color, m.meta, m.href, m.parent_id
        FROM menu m
        JOIN menu_tree mt ON m.parent_id = mt.id
        WHERE m.deleted = 0 AND m.status = 1
    )
    SELECT DISTINCT id, deleted, status, remark, label, path, component, icon, icon_bg_color, meta, href, parent_id
    FROM menu_tree
    `
	if menu.Remark.Remark != "" {
		query += fmt.Sprintf(" WHERE remark LIKE '%%%s%%'", menu.Remark.Remark)
	}
	var menuList []entity.Menu
	if err := config.Boot.Mysql.Raw(query).Scan(&menuList).Error; err != nil {
		return err
	}
	// 开启事务
	tx := config.Boot.Mysql.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 循环删除 menu 表和 menu_permission 表中的记录
	for _, m := range menuList {
		// 删除 menu 表中的记录
		if err := tx.Model(menu).Where("id = ?", m.Id.Id).Update("deleted", 1).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
		// 删除 menu_permission 表中的记录
		if err := tx.Model(&entity.MenuPermission{}).Where("id = ?", m.Id.Id).Update("deleted", 1).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
	}
	// 提交事务
	return tx.Commit().Error
}
