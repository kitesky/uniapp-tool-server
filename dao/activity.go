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

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type activity struct{}

func NewActivity() *activity {
	return &activity{}
}

func (s *activity) GetActivityWithUUID(uuid string) (log *models.UserActivityLog, err error) {
	log = &models.UserActivityLog{}
	if err = boot.DB.Model(&models.UserActivityLog{}).Where("uuid = ?", uuid).First(log).Error; err != nil {
		return
	}

	return
}

// 获取列表-分页
func (s *activity) GetActivityList(req schemas.ActivityPageReq) (result schemas.ActivityPageRes, err error) {
	list := []*models.UserActivityLog{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询字段信息
	fields := []string{"uuid", "code", "amount", "title", "description", "content_type", "status", "created_at", "updated_at"}

	// 查询条件
	query := boot.DB.Model(&models.UserActivityLog{}).Where("user_id = ?", req.UserID).Select(fields)

	switch req.Status {
	case "all":
		// 跳过查询全部
	default:
		query = query.Where("status = ?", req.Status)
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

	result = schemas.ActivityPageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}

func (s *activity) ActivityLog(tx *gorm.DB, log *models.UserActivityLog) (logRes *models.UserActivityLog, err error) {
	logRes = log

	// 开启事务
	err = tx.Transaction(func(tx *gorm.DB) error {
		// 创建日志
		if err = tx.Create(logRes).Error; err != nil {
			utils.ZapLog().Error("activity", "创建日志错误", zap.Error(err))
			return err
		}

		// 扣费规则:
		if logRes.Amount > 0 {
			// 获取用户信息
			user := models.User{}
			err = tx.Where("id = ?", logRes.UserID).Preload("VIP").First(&user).Error
			if err != nil {
				utils.ZapLog().Error("activity", "获取用户信息错误", zap.Error(err))
				return err
			}

			var amount float64 = 0
			var description string
			// // 检查是否VIP
			// if user.VIP.Active == "Y" && time.Now().Before(user.VIP.ExpireTime) {
			// 	// 扣费规则: 会员免费
			// 	amount = 0
			// 	description = "会员免费"
			// }

			if float64(user.Point) < logRes.Amount {
				// 点数不足
				utils.ZapLog().Error("activity", "点数不足", zap.Error(err))
				return errors.New("点数不足")
			}
			amount = logRes.Amount
			description = "点数消费"

			// 扣点并写入日志
			if _, err = NewPoint().PointLog(tx, &models.UserPointLog{
				UserID:      logRes.UserID,
				Code:        logRes.Code,
				Amount:      int64(amount),
				Type:        consts.PointActionDec,
				Title:       logRes.Title,
				Description: description,
			}); err != nil {
				utils.ZapLog().Error("activity", "扣费错误", zap.Error(err))
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})

	return
}
