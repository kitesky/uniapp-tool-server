package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
	"fmt"
	"math"
	"strings"
)

type tool struct{}

func NewTool() *tool {
	return &tool{}
}

func (o *tool) GetToolList(req schemas.ToolPageReq) (result schemas.ToolPageRes, err error) {
	list := []*models.AppTool{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.AppTool{})
	keyword := strings.TrimSpace(req.KeyWord)
	if keyword != "" {
		query.Where("title like ?", "%"+keyword+"%")
		NewSearch().SearchLog(&models.SearchLog{Title: keyword})
	}

	// 统计总数量
	if err = query.Count(&total).Error; err != nil {
		return
	}

	// 查询列表数据
	if err = query.Order(order).Offset(offsetSize).Limit(req.PageSize).Find(&list).Error; err != nil {
		return
	}

	// 计算总页数
	var totalPage int64 = 0
	if total > 0 {
		totalPage = int64(math.Ceil(float64(total) / float64(req.PageSize)))
	}

	result = schemas.ToolPageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (o *tool) GetToolWithCode(code string) (result *models.AppTool, err error) {
	result = &models.AppTool{}
	err = boot.DB.Where("code = ?", code).Find(&result).Error
	return
}

func (o *tool) GetTool(id int64) (result *models.AppTool, err error) {
	result = &models.AppTool{}
	err = boot.DB.Where("id = ?", id).Find(&result).Error
	return
}
