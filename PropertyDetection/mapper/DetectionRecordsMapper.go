package mapper

import (
	"PropertyDetection/config"
	"PropertyDetection/model/base"
	"PropertyDetection/model/entity"
)

type DetectionRecordsMapper struct {
}

func (mapper DetectionRecordsMapper) Page(page *base.Page, dr *entity.DetectionRecords) error {
	query := config.Boot.Mysql.Select("id, score, url, create_time, status, deleted").Where("deleted = ?", 0)
	var total int64 // 计算总数
	if err := config.Boot.Mysql.Model(dr).Count(&total).Error; err != nil {
		return err
	}
	var drs []entity.DetectionRecords
	if err := query.Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&drs).Error; err != nil {
		return err
	}
	page.Total = total
	page.Records = drs
	return nil
}
func (mapper DetectionRecordsMapper) RelationPatent(ids []interface{}) []entity.Patent {
	var patents []entity.Patent
	if err := config.Boot.Mysql.Model(entity.Patent{}).Where("id in ?", ids).Find(&patents).Error; err != nil {
		return []entity.Patent{}
	}
	return patents
}
