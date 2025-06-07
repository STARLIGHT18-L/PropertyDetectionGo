package api

import (
	"PropertyDetection/mapper"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ImageApi struct {
	imageMapper     mapper.ImageMapper
	patentMapper    mapper.PatentMapper
	copyrightMapper mapper.CopyrightMapper
	trademarkMapper mapper.TrademarkMapper
}

func (api ImageApi) Add(c *gin.Context) {
	object := map[string]interface{}{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	tabIndex, ok := object["tabIndex"].(string)
	if !ok {
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	var err error
	switch tabIndex {
	case "tab1":
		image := entity.Image{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &image)
		err = api.imageMapper.Insert(&image, c)
	case "tab2":
		trademark := entity.Trademark{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &trademark)
		err = api.trademarkMapper.Insert(&trademark, c)
	case "tab3":
		patent := entity.Patent{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &patent)
		err = api.patentMapper.Insert(&patent, c)
	case "tab4":
		copyright := entity.Copyright{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &copyright)
		err = api.copyrightMapper.Insert(&copyright, c)
	default:
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	if err != nil {
		c.JSON(400, tool.FailMsg("获取列表失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api ImageApi) Edit(c *gin.Context) {
	object := map[string]interface{}{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	tabIndex, ok := object["tabIndex"].(string)
	if !ok {
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	var err error
	switch tabIndex {
	case "tab1":
		image := entity.Image{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &image)
		err = api.imageMapper.Update(&image, c)
	case "tab2":
		trademark := entity.Trademark{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &trademark)
		err = api.trademarkMapper.Update(&trademark, c)
	case "tab3":
		patent := entity.Patent{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &patent)
		err = api.patentMapper.Update(&patent, c)
	case "tab4":
		copyright := entity.Copyright{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &copyright)
		err = api.copyrightMapper.Update(&copyright, c)
	default:
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	if err != nil {
		c.JSON(400, tool.FailMsg("获取列表失败"))
		return
	}
	c.JSON(200, tool.Success())
}
func (api ImageApi) Del(c *gin.Context) {
	object := map[string]interface{}{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	tabIndex, ok := object["tabIndex"].(string)
	if !ok {
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	var err error
	switch tabIndex {
	case "tab1":
		image := entity.Image{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &image)
		err = api.imageMapper.Delete(&image)
	case "tab2":
		trademark := entity.Trademark{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &trademark)
		err = api.trademarkMapper.Delete(&trademark)
	case "tab3":
		patent := entity.Patent{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &patent)
		err = api.patentMapper.Delete(&patent)
	case "tab4":
		copyright := entity.Copyright{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &copyright)
		err = api.copyrightMapper.Delete(&copyright)
	default:
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	if err != nil {
		c.JSON(400, tool.FailMsg("获取列表失败"))
		return
	}
	c.JSON(200, tool.Success())
}

func (api ImageApi) GetPage(c *gin.Context) {
	object := map[string]interface{}{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	page := base.Page{
		Current: int(object["current"].(float64)),
		Size:    int(object["size"].(float64)),
	}
	tabIndex, ok := object["tabIndex"].(string)
	if !ok {
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	var err error
	switch tabIndex {
	case "tab1":
		image := entity.Image{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &image)
		err = api.imageMapper.Page(&image, &page)
	case "tab2":
		trademark := entity.Trademark{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &trademark)
		err = api.trademarkMapper.Page(&trademark, &page)
	case "tab3":
		patent := entity.Patent{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &patent)
		err = api.patentMapper.Page(&patent, &page)
	case "tab4":
		copyright := entity.Copyright{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &copyright)
		err = api.copyrightMapper.Page(&copyright, &page)
	default:
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	if err != nil {
		c.JSON(400, tool.FailMsg("获取列表失败"))
		return
	}
	c.JSON(200, tool.SuccessData(page))
}
func (api ImageApi) Search(c *gin.Context) {
	object := map[string]interface{}{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	tabIndex, ok := object["tabIndex"].(string)
	if !ok {
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	var err error
	var results []entity.SearchResult
	switch tabIndex {
	case "tab1":
		image := entity.Image{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &image)
		results = api.imageMapper.Search(&image, c)
	case "tab2":
		trademark := entity.Trademark{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &trademark)
		results = api.trademarkMapper.Search(&trademark, c)
	case "tab3":
		patent := entity.Patent{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &patent)
		results = api.patentMapper.Search(&patent, c)
	case "tab4":
		copyright := entity.Copyright{}
		obj, _ := json.Marshal(object)
		err = json.Unmarshal(obj, &copyright)
		results = api.copyrightMapper.Search(&copyright, c)
	default:
		c.JSON(400, tool.FailMsg("请选择正确tab"))
		return
	}
	if err != nil {
		c.JSON(400, tool.FailMsg("获取列表失败"))
		return
	}
	c.JSON(200, tool.SuccessData(results))
}

func (api ImageApi) GetRelationPatent(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	var list, ok = json["ids"].([]interface{})
	if !ok {
		list = []interface{}{}
	}
	c.JSON(200, tool.SuccessData(api.patentMapper.GetRelationPatent(list)))
}
func (api ImageApi) SetRelationPatent(c *gin.Context) {
	image := entity.Image{}
	if err := c.ShouldBind(&image); err != nil {
		c.JSON(400, tool.FailMsg("参数格式错误"))
		return
	}
	err := api.imageMapper.SetRelationPatent(&image)
	if err != nil {
		c.JSON(500, tool.FailMsg("关联失败"))
	}
	c.JSON(200, tool.Success())
}
