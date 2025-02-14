package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"
	"fmt"
	"math"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type score struct{}

func NewScore() *score {
	return &score{}
}

// 积分兑换余额
func (s *score) ScoreExchange(req *schemas.ScoreExchangeReq) (err error) {
	err = boot.DB.Transaction(func(tx *gorm.DB) error {
		var min int64 = 100
		if req.Amount < min {
			return errors.New("最低兑换限制100积分")
		}

		// 查询用户积分
		user := &models.User{}
		if err = tx.Where("id = ?", req.UserID).First(user).Error; err != nil {
			return err
		}
		fmt.Println(111, user, req)
		// 积分不足
		if user.Score < req.Amount {
			return errors.New("积分不足")
		}

		// 计算积分可兑换的金额(分) - 除以100取整
		centTotal := decimal.NewFromInt(req.Amount).Div(decimal.NewFromInt(min)).RoundDown(0).IntPart()
		// 需要扣除的积分数量
		decScoreTotal := decimal.NewFromInt(centTotal).Mul(decimal.NewFromInt(min)).IntPart()
		// 分 转为 元
		incBalanceTotal := decimal.NewFromInt(centTotal).Div(decimal.NewFromInt(100)).RoundDown(2).InexactFloat64()

		// 写入积分日志和用户积分数量 - 减少积分
		if _, err = s.ScoreLog(tx, &models.UserScoreLog{
			UserID:      req.UserID,
			Amount:      int64(decScoreTotal),
			Type:        consts.ScoreActionDec,
			Code:        consts.ScoreActionExchange,
			Title:       "积分兑换现金",
			Description: "积分兑换现金余额",
		}); err != nil {
			return err
		}

		// 操作余额-增加余额
		if _, err = NewReward().RewardLog(tx, &models.UserRewardLog{
			UserID:      req.UserID,
			Amount:      incBalanceTotal,
			Type:        consts.BalanceActionInc,
			Title:       "积分兑换现金",
			Description: "积分兑换现金余额",
		}); err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return
}

// 获取订单列表-分页
func (s *score) GetScoreList(req schemas.ScorePageReq) (result schemas.ScorePageRes, err error) {
	list := []*models.UserScoreLog{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.UserScoreLog{}).Where("user_id = ?", req.UserID)

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

	result = schemas.ScorePageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (s *score) ScoreLog(tx *gorm.DB, log *models.UserScoreLog) (logRes *models.UserScoreLog, err error) {
	logRes = log

	// 开启事务
	err = tx.Transaction(func(tx *gorm.DB) error {
		// 创建日志
		if err = tx.Create(logRes).Error; err != nil {
			utils.ZapLog().Error("score", "创建日志错误", zap.Error(err))
			return err
		}

		// 更新用户积分
		expr := gorm.Expr("score + ?", logRes.Amount)
		switch log.Type {
		case consts.ScoreActionInc:
		case consts.ScoreActionDec:
			expr = gorm.Expr("score - ?", logRes.Amount)
		}

		if err = tx.Model(&models.User{}).Where("id = ?", logRes.UserID).UpdateColumn("score", expr).Error; err != nil {
			utils.ZapLog().Error("score", "变更用户积分信息错误", zap.Error(err))
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return
}
