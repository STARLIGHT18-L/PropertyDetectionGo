package mapper

import (
	"PropertyDetection/config"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
	"time"
)

type PatentMapper struct {
}

func (mapper PatentMapper) Insert(patent *entity.Patent, c *gin.Context) error {
	temp := entity.Patent{}
	config.Boot.Mysql.Model(&temp).Where("name = ? AND deleted = 0", patent.Name).First(&temp)
	if temp.Id.Id > 0 {
		return fmt.Errorf("图片已存在")
	}
	patent.Id.Id = entity.PrimaryKey()
	img := tool.DownloadImage(c.Request.Host, patent.Url, patent.Id.Id)
	finalGrayImg := tool.ConvertToGray(img)
	vectorImg := tool.ExtractFeatureVector(finalGrayImg)
	vectorImg = tool.SelectFeatures(vectorImg, 1024)
	patent.Vector = fmt.Sprintf("%v", vectorImg)
	return config.Boot.Mysql.Model(&patent).Create(&patent).Error
}
func (mapper PatentMapper) Update(patent *entity.Patent, c *gin.Context) error {
	img := tool.DownloadImage(c.Request.Host, patent.Url, patent.Id.Id)
	finalGrayImg := tool.ConvertToGray(img)
	vectorImg := tool.ExtractFeatureVector(finalGrayImg)
	vectorImg = tool.SelectFeatures(vectorImg, 1024)
	patent.Vector = fmt.Sprintf("%v", vectorImg)
	return config.Boot.Mysql.Model(&patent).Where("id = ?", patent.Id.Id).Updates(&patent).Error
}
func (mapper PatentMapper) Delete(patent *entity.Patent) error {
	return config.Boot.Mysql.Model(&patent).Where("id = ?", patent.Id.Id).Update("deleted", 1).Error
}

func (mapper PatentMapper) Page(patent *entity.Patent, page *base.Page) error {
	query := config.Boot.Mysql.Select("id", "name", "url", "content", "owner", "register_date", "status", "remark", "deleted").Where("deleted = ?", 0)
	if patent.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", patent.Name))
	}
	if patent.Content != "" {
		query = query.Where("content LIKE ?", fmt.Sprintf("%%%s%%", patent.Content))
	}
	if patent.Owner != "" {
		query = query.Where("owner LIKE ?", fmt.Sprintf("%%%s%%", patent.Owner))
	}
	var total int64
	if err := config.Boot.Mysql.Model(patent).Count(&total).Error; err != nil {
		return err
	}
	var patents []entity.Patent
	if err := query.Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&patents).Error; err != nil {
		return err
	}
	page.Total = total
	page.Records = patents
	return nil
}
func (mapper PatentMapper) Search(patent *entity.Patent, c *gin.Context) []entity.SearchResult {
	img := tool.DownloadImage(c.Request.Host, patent.Url, entity.PrimaryKey())
	finalGrayImg := tool.ConvertToGray(img)
	vectorImg := tool.ExtractFeatureVector(finalGrayImg)
	vectorImg = tool.SelectFeatures(vectorImg, 1024)
	// 提前计算查询向量的模
	queryVectorNorm := 0.0
	for _, v := range vectorImg {
		queryVectorNorm += v * v
	}
	queryVectorNorm = math.Sqrt(queryVectorNorm)
	// 构建 SQL 查询
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT id, url, name, content, relation_patent, ")
	sqlBuilder.WriteString("( ")
	// 计算点积
	for i, v := range vectorImg {
		if i > 0 {
			sqlBuilder.WriteString(" + ")
		}
		sqlBuilder.WriteString(fmt.Sprintf("CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(vector, ' ', %d), ' ', -1) AS DECIMAL(10, 6)) * %f", i+1, v))
	}
	sqlBuilder.WriteString(" ) / ( ")
	// 计算数据库中向量的模
	sqlBuilder.WriteString("SQRT( ")
	for i := range vectorImg {
		if i > 0 {
			sqlBuilder.WriteString(" + ")
		}
		sqlBuilder.WriteString(fmt.Sprintf("POW(CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(vector, ' ', %d), ' ', -1) AS DECIMAL(10, 6)), 2)", i+1))
	}
	sqlBuilder.WriteString(" ) * ")
	// 乘以查询向量的模
	sqlBuilder.WriteString(strconv.FormatFloat(queryVectorNorm, 'f', 6, 64))
	sqlBuilder.WriteString(" ) AS score1 ")
	sqlBuilder.WriteString("FROM patent ")
	sqlBuilder.WriteString("WHERE deleted = 0 ")
	sqlBuilder.WriteString("ORDER BY score1 DESC ")
	sqlBuilder.WriteString(fmt.Sprintf("LIMIT %d", 15))
	sqlQuery := sqlBuilder.String()
	querys := make([]entity.SearchResult, 0, 15)
	results := make([]entity.SearchResult, 0, 15)
	// 执行 SQL 查询
	config.Boot.Mysql.Raw(sqlQuery).Scan(&querys)
	for _, result := range querys {
		// 获取 SQL 计算出的 score1
		score1, _ := strconv.ParseFloat(result.Score1, 64) // 假设 SearchResult 结构体有 Score1Float 方法获取 score1 的 float64 值
		// 格式化 score1 并添加百分号
		result.Score1 = fmt.Sprintf("%.2f%%", score1*100)
		// 计算 score2
		score2 := tool.CosineSimilarity(result.Content, patent.Content, []string{result.Content, patent.Content})
		result.Score2 = fmt.Sprintf("%.2f%%", score2*100)
		// 计算 score3
		score3 := tool.CosineSimilarity(result.Name, patent.Name, []string{result.Name, patent.Name})
		result.Score3 = fmt.Sprintf("%.2f%%", score3*100)
		// 计算最终得分
		finalScore := (score1*0.5 + score2*0.25 + score3*0.25) * 100
		result.Score = fmt.Sprintf("%.2f%%", finalScore)
		results = append(results, result)
	}
	if len(results) > 0 {
		dr := entity.DetectionRecords{
			Id:         entity.Id{Id: entity.PrimaryKey()},
			Status:     entity.Status{Status: 1},
			Deleted:    entity.Deleted{Deleted: 0},
			Url:        patent.Url,
			Score:      results[0].Score,
			CreateTime: time.Now(),
		}
		config.Boot.Mysql.Create(&dr)
	}
	return results
}
func (mapper PatentMapper) GetRelationPatent(ids []interface{}) map[string]interface{} {
	allSelections := []entity.Patent{}
	config.Boot.Mysql.Select("id,name").Where("deleted = ?", 0).Find(&allSelections)
	hadSelections := []entity.Patent{}
	config.Boot.Mysql.Where("deleted = ?", 0).Where("id IN (?)", ids).Find(&hadSelections)
	result := map[string]interface{}{
		"allSelections": allSelections,
		"hadSelections": hadSelections,
	}
	return result
}
