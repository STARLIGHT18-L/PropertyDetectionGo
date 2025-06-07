package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(r *gin.Engine) {
	router := r.Group("/api/role").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.POST("/rowSave", api.RoleApi{}.RowSave)
		router.POST("/rowEdit", api.RoleApi{}.RowEdit)
		router.POST("/rowDel", api.RoleApi{}.RowDel)
		router.GET("/getPage", api.RoleApi{}.GetPage)
		router.GET("/getRoleListMaps", api.RoleApi{}.GetRoleListMaps)
		router.GET("/getMenuTreeData", api.RoleApi{}.GetMenuTreeData)
		router.POST("/setMenuPermission", api.RoleApi{}.SetMenuPermission)
		router.GET("/getMenuTreeDataById", api.RoleApi{}.GetMenuTreeDataById)
	}
}
