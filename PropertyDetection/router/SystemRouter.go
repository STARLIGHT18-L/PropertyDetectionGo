package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(r *gin.Engine) {
	router := r.Group("/api/system").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.GET("/gc", api.SystemApi{}.Gc)
		router.GET("/cleanFile", api.SystemApi{}.CleanFile)
		router.POST("/upload", api.SystemApi{}.Upload)
	}
}
