package router

import "github.com/gin-gonic/gin"

func InitRouterGroup(router *gin.Engine) {
	//// 前端项目静态资源
	//router.Static("/assets", "./static/dist/assets")
	//router.StaticFile("/", "./static/dist/index.html")
	//router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	//// 其他静态资源
	//router.Static("/public", "./static")
	//router.Static("/storage", "./storage/app/public")
	InitAuthRouter(router)             //登录路由分组
	InitUserRouter(router)             //用户路由分组
	InitRoleRouter(router)             //角色路由分组
	InitMenuRouter(router)             //菜单路由分组
	InitSystemRouter(router)           //系统路由分组
	InitImageRouter(router)            //图片路由分组
	InitDetectionRecordsRouter(router) //检测记录路由分组
}
