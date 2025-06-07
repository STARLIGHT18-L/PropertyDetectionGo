package api

import (
	"PropertyDetection/mapper"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"github.com/gin-gonic/gin"
)

type DetectionRecordsApi struct {
	mapper mapper.DetectionRecordsMapper
}

func (api DetectionRecordsApi) GetPage(c *gin.Context) {
	page := base.Page{}
	if err := c.ShouldBind(&page); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	dr := entity.DetectionRecords{}
	if err := c.ShouldBind(&dr); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.mapper.Page(&page, &dr)
	if err != nil {
		c.JSON(400, tool.FailMsg("获取检测记录列表失败"))
		return
	}
	c.JSON(200, tool.SuccessData(page))
}
func (api DetectionRecordsApi) RelationPatent(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	var list, ok = json["ids"].([]interface{})
	if !ok {
		list = []interface{}{}
	}
	c.JSON(200, tool.SuccessData(api.mapper.RelationPatent(list)))
}
