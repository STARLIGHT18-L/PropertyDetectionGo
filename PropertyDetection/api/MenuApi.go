package api

import (
	"PropertyDetection/mapper"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	mapper mapper.MenuMapper
}

func (api MenuApi) GetMenuList(c *gin.Context) {
	menu := &entity.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	menuList, err := api.mapper.GetMenuList(menu)
	if err != nil {
		c.JSON(400, tool.FailMsg("获取菜单列表失败"))
		return
	}
	c.JSON(200, tool.SuccessData(menuList))
}
func (api MenuApi) GetMenuLabelAndId(c *gin.Context) {
	menuList, _ := api.mapper.GetMenuList(&entity.Menu{})
	menu := entity.Menu{
		Id: entity.Id{
			Id: 0,
		},
		Label:    "顶层菜单",
		Path:     "/",
		Icon:     "el-icon-document",
		Meta:     "{\"i18n\":\"test\"}",
		ParentId: 0,
		Children: []entity.Menu{},
	}
	menuList = append([]entity.Menu{menu}, menuList...)
	c.JSON(200, tool.SuccessData(menuList))
}
func (api MenuApi) RowSave(c *gin.Context) {
	menu := &entity.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	menu.Id.Id = entity.PrimaryKey()
	menu.Status.Status = 1
	menu.Deleted.Deleted = 0
	menuPermission := entity.MenuPermission{
		Id:      entity.Id{Id: entity.PrimaryKey()},
		Status:  entity.Status{Status: 1},
		Remark:  entity.Remark{Remark: "admin"},
		Deleted: entity.Deleted{Deleted: 0},
		MenuId:  menu.Id.Id,
		RoleId:  1,
	}
	if err := api.mapper.Insert(menu, &menuPermission); err != nil {
		c.JSON(400, tool.FailMsg("保存失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api MenuApi) RowEdit(c *gin.Context) {
	menu := &entity.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	if err := api.mapper.Update(menu); err != nil {
		c.JSON(400, tool.FailMsg("修改失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api MenuApi) RowDel(c *gin.Context) {
	menu := &entity.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	if err := api.mapper.Delete(menu); err != nil {
		c.JSON(400, tool.FailMsg("删除失败"))
		return
	}
	c.JSON(200, tool.Success())
}
