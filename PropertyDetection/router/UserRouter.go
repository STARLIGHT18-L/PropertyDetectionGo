package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(r *gin.Engine) {
	router := r.Group("/api/user")
	{
		router.GET("/getCaptcha", api.UserApi{}.GetCaptcha)
		router.POST("/login", api.UserApi{}.Login)
	}
}
func InitUserRouter(r *gin.Engine) {
	router := r.Group("/api/user").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.GET("/getUserInfo", api.UserApi{}.GetUserInfo)
		router.GET("/getTopMenu", api.UserApi{}.GetTopMenu)
		router.GET("/getMenu", api.UserApi{}.GetMenu)
		router.GET("/logout", api.UserApi{}.Logout)
		router.POST("/rowSave", api.UserApi{}.RowSave)
		router.POST("/rowEdit", api.UserApi{}.RowEdit)
		router.POST("/rowDel", api.UserApi{}.RowDel)
		router.GET("/getPage", api.UserApi{}.GetPage)
	}
}
