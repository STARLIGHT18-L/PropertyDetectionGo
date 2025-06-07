package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(r *gin.Engine) {
	router := r.Group("/api/menu").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.GET("/getMenuList", api.MenuApi{}.GetMenuList)
		router.GET("/getMenuLabelAndId", api.MenuApi{}.GetMenuLabelAndId)
		router.POST("/rowSave", api.MenuApi{}.RowSave)
		router.POST("/rowEdit", api.MenuApi{}.RowEdit)
		router.POST("/rowDel", api.MenuApi{}.RowDel)
	}
}
