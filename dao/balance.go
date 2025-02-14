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

type balance struct{}

func NewBalance() *balance {
	return &balance{}
}

// 获取订单列表-分页
func (o *balance) GetBalanceList(req schemas.BalancePageReq) (result schemas.BalancePageRes, err error) {
	list := []*models.UserBalanceLog{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.UserBalanceLog{}).Where("user_id = ?", req.UserID)

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

	result = schemas.BalancePageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (s *balance) BalanceLog(tx *gorm.DB, log *models.UserBalanceLog) (newLog *models.UserBalanceLog, err error) {
	newLog = log

	user := models.User{}
	if err = tx.Where("id", newLog.UserID).Find(&user).Error; err != nil {
		utils.ZapLog().Error("balance", "查询用户信息错误", zap.Error(err))
		return
	}

	// 创建日志
	newLog.Balance = user.Balance // 当前余额
	if err = tx.Create(newLog).Error; err != nil {
		utils.ZapLog().Error("balance", "创建日志错误", zap.Error(err))
		return
	}

	// 更新用户余额
	expr := gorm.Expr("balance + ?", newLog.Amount)
	switch log.Type {
	case consts.BalanceActionInc:
	case consts.BalanceActionDec:
		expr = gorm.Expr("balance - ?", newLog.Amount)
	}

	if err = tx.Model(&models.User{}).Where("id = ?", newLog.UserID).UpdateColumn("balance", expr).Error; err != nil {
		utils.ZapLog().Error("balance", "变更用户余额信息错误", zap.Error(err))
		return
	}

	return
}
