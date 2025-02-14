package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type reward struct{}

func NewReward() *reward {
	return &reward{}
}

// 奖励金提现
func (s *reward) RewardExchange(req *schemas.RewardExchangeReq) (err error) {
	var min float64 = 0.3
	if req.Amount < min {
		return errors.New("最低提现金额限制0.3元")
	}

	tx := boot.DB.Begin()

	// 查询用户现金
	user := models.User{}
	if err = tx.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 现金不足
	if user.Reward < req.Amount {
		tx.Rollback()
		return errors.New("金额不足")
	}

	// 写入现金日志 - 减少现金
	if _, err = s.RewardLog(tx, &models.UserRewardLog{
		UserID:      req.UserID,
		Amount:      req.Amount, // 根据提现金额扣除
		Type:        consts.RewardActionDec,
		Code:        consts.RewardActionExchange,
		Title:       "奖励金提现",
		Description: "奖励金提现",
	}); err != nil {
		tx.Rollback()
		return err
	}

	// 操作余额-增加余额
	if _, err = NewUserTransfer().CreateUserTransferOrder(tx, req.UserID, req.Amount, req.PaymentID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 奖励金兑换点数
func (s *reward) RewardExchangePoint(req *schemas.RewardExchangeReq) (err error) {
	tx := boot.DB.Begin()

	// 查询用户现金
	user := models.User{}
	if err = tx.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 查询产品详情
	product, err := NewProduct().GetProduct(req.ProductID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 现金不足
	if user.Reward < product.Price {
		tx.Rollback()
		return errors.New("金额不足")
	}

	// 写入现金日志 - 减少现金
	if _, err = s.RewardLog(tx, &models.UserRewardLog{
		UserID:      req.UserID,
		Amount:      product.Price, // 根据产品价格扣除
		Type:        consts.RewardActionDec,
		Code:        consts.RewardActionExchange,
		Title:       "奖励金兑换点数",
		Description: "奖励金兑换点数",
	}); err != nil {
		tx.Rollback()
		return err
	}

	// 产品 extra
	extra := schemas.ProductPointExtra{}
	err = json.Unmarshal([]byte(product.Extra), &extra)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入点数日志&更新用户点数
	if _, err := NewPoint().PointLog(tx, &models.UserPointLog{
		UserID:      req.UserID,
		Code:        product.Code,
		Amount:      extra.Point + extra.Gift,
		Title:       product.Title,
		Description: product.Description,
		Type:        consts.PointActionInc,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 获取订单列表-分页
func (s *reward) GetRewardList(req schemas.RewardPageReq) (result schemas.RewardPageRes, err error) {
	list := []*models.UserRewardLog{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.UserRewardLog{}).Where("user_id = ?", req.UserID)

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

	result = schemas.RewardPageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (s *reward) RewardLog(tx *gorm.DB, log *models.UserRewardLog) (newLog *models.UserRewardLog, err error) {
	// 创建日志
	if err = tx.Create(log).Error; err != nil {
		utils.ZapLog().Error("reward", "创建日志错误", zap.Error(err))
		return
	}

	// 更新
	expr := gorm.Expr("reward + ?", log.Amount)
	switch log.Type {
	case consts.RewardActionInc:
	case consts.RewardActionDec:
		expr = gorm.Expr("reward - ?", log.Amount)
	}

	if err = tx.Model(&models.User{}).Where("id = ?", log.UserID).UpdateColumn("reward", expr).Error; err != nil {
		utils.ZapLog().Error("reward", "变更信息错误", zap.Error(err))
		return
	}

	return log, nil
}
