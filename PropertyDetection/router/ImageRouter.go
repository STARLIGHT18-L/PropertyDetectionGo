package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitImageRouter(r *gin.Engine) {
	router := r.Group("/api/image").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.POST("/add", api.ImageApi{}.Add)
		router.POST("/edit", api.ImageApi{}.Edit)
		router.POST("/del", api.ImageApi{}.Del)
		router.POST("/getPage", api.ImageApi{}.GetPage)
		router.POST("/search", api.ImageApi{}.Search)
		router.POST("/getRelationPatent", api.ImageApi{}.GetRelationPatent)
		router.POST("/setRelationPatent", api.ImageApi{}.SetRelationPatent)
	}
}
