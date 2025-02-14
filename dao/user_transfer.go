package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"
	"fmt"
	"math"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userTransfer struct{}

func NewUserTransfer() *userTransfer {
	return &userTransfer{}
}

func (u *userTransfer) CreateUserTransferOrder(tx *gorm.DB, userID int64, amount float64, paymentID int64) (result *models.UserTransfer, err error) {
	// 查询收款信息
	payment := models.UserPaymentAccount{}
	tx.Where("user_id = ? and id = ?", userID, paymentID).First(&payment)
	if payment.ID == 0 {
		utils.ZapLog().Error("reward", "收款信息错误")
		return nil, errors.New("收款信息错误")
	}

	result = &models.UserTransfer{
		UserID:  userID,
		Amount:  amount,
		OrderNo: utils.GenerateStringUniqueID(),
		PayType: payment.PayType,
		Account: payment.Account,
		Name:    payment.Name,
		Status:  "pending",
	}

	if err = tx.Create(result).Error; err != nil {
		utils.ZapLog().Error("reward", "创建订单错误", zap.Error(err))
		return
	}

	return
}

// 获取列表-分页
func (s *userTransfer) GetUserTransferList(req schemas.UserTransferPageReq) (result schemas.UserTransferPageRes, err error) {
	list := []*models.UserTransfer{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.UserTransfer{}).Where("user_id = ?", req.UserID)

	// switch req.Status {
	// case "all":
	// 	// 跳过查询全部
	// default:
	// 	query = query.Where("status = ?", req.Status)
	// }

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

	result = schemas.UserTransferPageRes{
		Items:     list,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}
