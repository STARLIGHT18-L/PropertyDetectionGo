package router

import (
	"PropertyDetection/api"
	"PropertyDetection/config"
	"PropertyDetection/middleware"
	"github.com/gin-gonic/gin"
)

func InitDetectionRecordsRouter(r *gin.Engine) {
	router := r.Group("/api/detectionRecords").Use(middleware.JwtAuth(config.AppGuardName))
	{
		router.GET("/getPage", api.DetectionRecordsApi{}.GetPage)
		router.POST("/relationPatent", api.DetectionRecordsApi{}.RelationPatent)
	}
}
