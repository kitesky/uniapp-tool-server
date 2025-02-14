package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"fmt"
	"math"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type point struct{}

func NewPoint() *point {
	return &point{}
}

// 获取列表-分页
func (s *point) GetPointList(req schemas.PointPageReq) (result schemas.PointPageRes, err error) {
	list := []*models.UserPointLog{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.UserPointLog{}).Where("user_id = ?", req.UserID)

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

	result = schemas.PointPageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (s *point) PointLog(tx *gorm.DB, log *models.UserPointLog) (logRes *models.UserPointLog, err error) {
	logRes = log

	// 开启事务
	err = tx.Transaction(func(tx *gorm.DB) error {
		// 创建日志
		if err = tx.Create(logRes).Error; err != nil {
			utils.ZapLog().Error("point", "创建日志错误", zap.Error(err))
			return err
		}

		// 更新
		expr := gorm.Expr("point + ?", logRes.Amount)
		switch log.Type {
		case consts.PointActionInc:
		case consts.PointActionDec:
			expr = gorm.Expr("point - ?", logRes.Amount)
		}

		if err = tx.Model(&models.User{}).Where("id = ?", logRes.UserID).UpdateColumn("point", expr).Error; err != nil {
			utils.ZapLog().Error("point", "变更信息错误", zap.Error(err))
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return
}
