package api

import (
	"PropertyDetection/config"
	"PropertyDetection/mapper"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type RoleApi struct {
	mapper     mapper.RoleMapper
	userMapper mapper.UserMapper
}

func (api RoleApi) RowSave(c *gin.Context) {
	role := entity.Role{}
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	role.Id.Id = entity.PrimaryKey()
	role.Status.Status = 1
	role.Deleted.Deleted = 0
	err := api.mapper.Insert(&role)
	if err != nil {
		c.JSON(400, tool.FailMsg("添加角色失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api RoleApi) RowEdit(c *gin.Context) {
	role := entity.Role{}
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.mapper.Update(&role)
	if err != nil {
		c.JSON(400, tool.FailMsg("修改角色失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api RoleApi) RowDel(c *gin.Context) {
	role := entity.Role{}
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	if err := api.mapper.Delete(&role); err != nil {
		c.JSON(400, tool.FailMsg("删除角色失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api RoleApi) GetPage(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	userStr := config.Boot.Cache.GetValue(tokenStr)
	info := entity.User{}
	userByte, ok := userStr.([]byte)
	if !ok {
		c.JSON(400, tool.FailMsg("用户信息获取失败"))
		return
	}
	json.Unmarshal(userByte, &info)
	page := base.Page{}
	if err := c.ShouldBind(&page); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	role := entity.Role{}
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.mapper.Page(&page, &role, &info)
	if err != nil {
		c.JSON(200, tool.SuccessData([]interface{}{}))
		return
	}
	c.JSON(200, tool.SuccessData(page))
}
func (api RoleApi) GetMenuTreeData(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	userStr := config.Boot.Cache.GetValue(tokenStr)
	info := entity.User{}
	userByte, ok := userStr.([]byte)
	if !ok {
		c.JSON(400, tool.FailMsg("用户信息获取失败"))
		return
	}
	json.Unmarshal(userByte, &info)
	menuList, err := api.userMapper.GetMenuListByRoleId(info.RoleId)
	if err != nil {
		c.JSON(400, tool.SuccessData([]interface{}{}))
		return
	}
	c.JSON(200, tool.SuccessData(menuList))
}
func (api RoleApi) GetMenuTreeDataById(c *gin.Context) {
	roleId, _ := strconv.ParseInt(c.Query("roleId"), 10, 64)
	menuList, err := api.userMapper.GetMenuListByRoleId(int(roleId))
	if err != nil {
		c.JSON(400, tool.SuccessData([]interface{}{}))
		return
	}
	c.JSON(200, tool.SuccessData(menuList))
}
func (api RoleApi) GetRoleListMaps(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	userStr := config.Boot.Cache.GetValue(tokenStr)
	info := entity.User{}
	userByte, ok := userStr.([]byte)
	if !ok {
		c.JSON(400, tool.FailMsg("用户信息获取失败"))
		return
	}
	json.Unmarshal(userByte, &info)
	c.JSON(200, tool.SuccessData(api.mapper.ListMaps(&info)))
}
func (api RoleApi) SetMenuPermission(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	// 检查必要参数
	if json["roleName"] == nil || strings.TrimSpace(json["roleName"].(string)) == "" || json["roleId"] == nil || json["ids"] == nil {
		c.JSON(400, tool.FailMsg("参数错误"))
		return
	}
	roleId, _ := strconv.Atoi(fmt.Sprintf("%v", json["roleId"]))
	roleName := json["roleName"].(string)
	ids, ok := json["ids"].([]interface{})
	if !ok {
		c.JSON(400, tool.FailMsg("参数错误"))
		return
	}
	err := api.mapper.SetMenuPermission(roleId, roleName, ids)
	if err.Error != nil {
		c.JSON(400, tool.FailMsg("设置权限失败"))
		return
	}
	c.JSON(200, tool.Success())
}
