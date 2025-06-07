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
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

// UserApi 结构体用于处理用户相关的 API 请求
type UserApi struct {
	mapper mapper.UserMapper
}

// GetCaptcha 生成验证码
func (api UserApi) GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0.2, 10)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	key, image, err := captcha.Generate()
	if err != nil {
		c.JSON(400, tool.FailMsg("验证码生成失败"))
		return
	}
	value := base64Captcha.DefaultMemStore.Get(key, true)
	config.Boot.Cache.SetValueExpiry(key, value, time.Minute*5)
	c.JSON(200, tool.SuccessMsgData("验证码获取成功", map[string]interface{}{
		"key":   key,
		"image": image,
	}))
}

// Login 用户登录
func (api UserApi) Login(c *gin.Context) {
	var obj = make(map[string]interface{})
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	user := entity.User{
		Username: fmt.Sprint(obj["username"]),
	}
	code := obj["code"]
	redomStr := obj["redomStr"]
	value := config.Boot.Cache.GetValue(fmt.Sprint(redomStr))
	if value != code {
		c.JSON(400, tool.FailMsg("验证码错误"))
		return
	}
	user.Username, _ = tool.Decrypt(user.Username)
	temp, err := api.mapper.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(400, tool.FailMsg("用户不存在"))
		return
	}
	// 直接使用 obj["password"] 作为 map[string]interface{}
	passwordJSON, ok := obj["password"].(map[string]interface{})
	if !ok {
		c.JSON(400, tool.FailMsg("密码格式错误"))
		return
	}
	wordsInterface, ok := passwordJSON["words"] // 提取 words 数组
	if !ok {
		c.JSON(400, tool.FailMsg("密码格式错误，缺少 'words' 字段"))
		return
	}
	wordsJSON, err := json.Marshal(wordsInterface)
	if err != nil {
		c.JSON(400, tool.FailMsg("密码解析错误"))
		return
	}
	var words []int
	err = json.Unmarshal(wordsJSON, &words)
	if err != nil {
		c.JSON(400, tool.FailMsg("密码解析错误，无法将 words 转换为整数切片"))
		return
	}
	if !tool.Sha256Decrypt(words, temp.Password) {
		c.JSON(400, tool.FailMsg("密码错误"))
		return
	}
	token, err, _ := config.JwtService.CreateToken(config.AppGuardName, temp)
	if err != nil {
		c.JSON(400, tool.FailMsg("登录失败"))
		return
	}
	userJSON, _ := json.Marshal(temp)
	config.Boot.Cache.SetValueExpiry(token.AccessToken,
		userJSON,
		time.Duration(config.Boot.Config.Jwt.JwtTtl*1000000000))
	c.JSON(200, tool.SuccessData(token.AccessToken))
}

// GetUserInfo 获取用户信息
func (api UserApi) GetUserInfo(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	userStr := config.Boot.Cache.GetValue(tokenStr)
	user := entity.User{}
	userByte, ok := userStr.([]byte)
	if !ok {
		c.JSON(400, tool.FailMsg("用户信息获取失败"))
		return
	}
	json.Unmarshal(userByte, &user)
	user.Password = ""                   // 清空用户密码
	data := make(map[string]interface{}) // 构建返回数据
	data["userInfo"] = user
	roleNames, _ := api.mapper.GetRoleNamesByRoleId(user.RoleId)
	data["roles"] = roleNames
	// 定义权限列表
	permissions := []string{
		"sys_crud_btn_add",
		"sys_crud_btn_export",
		"sys_menu_btn_add",
		"sys_menu_btn_edit",
		"sys_menu_btn_del",
		"sys_role_btn1",
		"sys_role_btn2",
		"sys_role_btn3",
		"sys_role_btn4",
		"sys_role_btn5",
		"sys_role_btn6"}
	data["permission"] = permissions
	fmt.Printf("用户获取个人信息==> %v\n", data)
	c.JSON(200, tool.SuccessData(data))
}
func (api UserApi) GetMenu(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	userStr := config.Boot.Cache.GetValue(tokenStr)
	user := entity.User{}
	userByte, ok := userStr.([]byte)
	if !ok {
		c.JSON(400, tool.FailMsg("用户信息获取失败"))
		return
	}
	json.Unmarshal(userByte, &user)
	menus, _ := api.mapper.GetMenuListByRoleId(user.RoleId)
	c.JSON(200, tool.SuccessData(menus))
}
func (api UserApi) GetTopMenu(c *gin.Context) {
	menus := []entity.Menu{
		{
			Label:       "首页",
			Path:        "/wel/index",
			Icon:        "el-icon-document",
			Meta:        "{\"i18n\":\"dashboard\"}",
			ParentId:    0,
			IconBgColor: "#fff",
			Children:    []entity.Menu{},
		},
		{
			Label:       "测试",
			Path:        "/test",
			Icon:        "el-icon-document",
			Meta:        "{\"i18n\":\"test\"}",
			ParentId:    0,
			IconBgColor: "#fff",
			Children:    []entity.Menu{},
		},
		{
			Label:       "更多",
			Path:        "/wel/more",
			Icon:        "el-icon-document",
			Meta:        "{\"menu\":false,\"i18n\":\"more\"}",
			ParentId:    0,
			IconBgColor: "#fff",
			Children:    []entity.Menu{},
		},
	}
	c.JSON(200, tool.SuccessData(menus))
}
func (api UserApi) Logout(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.TrimSpace(strings.SplitN(tokenStr, " ", 2)[1])
	config.Boot.Cache.DeleteValue(tokenStr)
	c.JSON(200, tool.Success())
}
func (api UserApi) RowSave(c *gin.Context) {
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	user.Password, _ = tool.Decrypt(user.Password)
	user.Id.Id = entity.PrimaryKey()
	user.Status.Status = 1
	user.Deleted.Deleted = 0
	err := api.mapper.Insert(&user)
	if err != nil {
		c.JSON(400, tool.FailMsg("添加用户失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api UserApi) RowEdit(c *gin.Context) {
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	user.Password = ""
	err := api.mapper.Update(&user)
	if err != nil {
		c.JSON(400, tool.FailMsg("修改用户失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api UserApi) RowDel(c *gin.Context) {
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.mapper.Delete(&user)
	if err != nil {
		c.JSON(400, tool.FailMsg("删除用户失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api UserApi) GetPage(c *gin.Context) {
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
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.mapper.Page(&page, &user, &info)
	if err != nil {
		c.JSON(200, tool.SuccessData([]interface{}{}))
		return
	}
	c.JSON(200, tool.SuccessData(page))
}
