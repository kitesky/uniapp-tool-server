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
	"time"

	"gorm.io/gorm"
)

type order struct{}

func NewOrder() *order {
	return &order{}
}

func (s *order) GetOrder(orderNo string) (order *models.Order, err error) {
	order = &models.Order{}
	err = boot.DB.Where("order_no = ?", orderNo).First(order).Error
	return
}

func (s *order) ProductOrderProcess(orderNo string, payOrderNo string, payTime time.Time) (order *models.Order, err error) {
	err = boot.DB.Transaction(func(tx *gorm.DB) error {
		order = &models.Order{}

		// 查询订单数据
		if err := tx.Where("order_no = ? and status = ?", orderNo, "pending").First(order).Error; err != nil {
			// 查询订单错误回滚
			return err
		}

		// 更新订单状态
		if err := tx.Model(order).Where("order_no", orderNo).Update("status", "completed").Error; err != nil {
			// 更新订单状态错误回滚
			return err
		}

		// 查询产品详情
		product, err := NewProduct().GetProduct(order.ProductID)
		if err != nil {
			return err
		}

		switch product.Type {
		case "vip":
			extra := schemas.ProductVipExtra{}
			err := json.Unmarshal([]byte(order.Extra), &extra)
			if err != nil {
				return err
			}

			// 更新会员状态 过期时间
			if _, err := NewUser().UpdateUserVIP(order.UserID, extra.Month); err != nil {
				return err
			}

			// 赠送积分
			if extra.Point > 0 {
				// 写入点数日志&更新用户点数
				if _, err := NewPoint().PointLog(tx, &models.UserPointLog{
					UserID:      order.UserID,
					Code:        product.Code,
					Amount:      extra.Point,
					Title:       product.Title,
					Description: product.Description,
					Type:        consts.PointActionInc,
				}); err != nil {
					return err
				}
			}
		case "point":
			extra := schemas.ProductPointExtra{}
			err := json.Unmarshal([]byte(order.Extra), &extra)
			if err != nil {
				return err
			}

			// 写入点数日志&更新用户点数
			if _, err := NewPoint().PointLog(tx, &models.UserPointLog{
				UserID:      order.UserID,
				Code:        product.Code,
				Amount:      extra.Point + extra.Gift,
				Title:       product.Title,
				Description: product.Description,
				Type:        consts.PointActionInc,
			}); err != nil {
				return err
			}
		default:
			return errors.New("产品类型错误")
		}

		return nil
	})

	return
}

func (s *order) RechargeOrderProcess(orderNo string, payOrderNo string, payTime time.Time) (order *models.Order, err error) {
	err = boot.DB.Transaction(func(tx *gorm.DB) error {
		order = &models.Order{}

		// 查询订单数据
		if err := tx.Where("order_no = ? and status = ? and order_type = ?", orderNo, "pending", "recharge").First(order).Error; err != nil {
			// 查询订单错误回滚
			return err
		}

		// 更新订单状态
		if err := tx.Model(order).Where("order_no", orderNo).Update("status", "completed").Error; err != nil {
			// 更新订单状态错误回滚
			return err
		}

		// 更新账户余额和日志
		if _, err := NewBalance().BalanceLog(tx, &models.UserBalanceLog{
			UserID:      order.UserID,
			Amount:      order.Amount,
			Title:       "余额充值",
			Description: fmt.Sprintf("充值订单-%s", order.OrderNo),
			Type:        "increase",
		}); err != nil {
			return err
		}

		return nil
	})

	return
}

func (s *order) CreateOrder(req *schemas.OrderCreateReq) (newOrder *models.Order, err error) {
	newOrder = &models.Order{
		UserID:      req.UserID,
		OrderNo:     utils.GenerateStringUniqueID(),
		OrderType:   req.OrderType,
		PayType:     req.PayType,
		Title:       schemas.OrderTypeOptions[req.OrderType].Title,
		Description: schemas.OrderTypeOptions[req.OrderType].Description,
		ExpiredAt:   time.Now().Add(consts.OrderExpireTime * time.Minute),
	}

	// 订单类型
	switch req.OrderType {
	case "vip", "point", "product":
		product, err := NewProduct().GetProduct(req.ProductID)
		if err != nil {
			return nil, err
		}

		newOrder.ProductID = product.ID
		newOrder.ProductCode = product.Code
		newOrder.ProductName = product.Name
		newOrder.ProductPrice = product.Price
		newOrder.Title = product.Title
		newOrder.Description = product.Description
		newOrder.Extra = product.Extra
		if req.Quantity == 0 {
			req.Quantity = 1
		}

		// 计算总价
		newOrder.Quantity = req.Quantity
		newOrder.Amount = product.Price * float64(req.Quantity)
	case "recharge":
		newOrder.Amount = req.Amount
	}

	if newOrder.Amount <= 0 {
		err = errors.New("订单金额错误")
		return
	}

	newOrder.PayAmount = newOrder.Amount
	err = boot.DB.Create(newOrder).Error
	return
}

// 获取订单列表-分页
func (s *order) GetOrderList(req schemas.OrderPageReq) (orderPage schemas.OrderPageRes, err error) {
	orderList := []*models.Order{}
	offsetSize := (req.Page - 1) * req.PageSize
	order := fmt.Sprintf("%s %s", req.SortField, req.SortType)
	var total int64 = 0

	// 查询条件
	query := boot.DB.Model(&models.Order{}).Where("user_id = ?", req.UserID)

	switch req.OrderStatus {
	case "all":
		// 跳过查询全部
	default:
		if req.OrderStatus != "" {
			query = query.Where("status = ?", req.OrderStatus)
		}
	}

	// 统计总数量
	if err = query.Count(&total).Error; err != nil {
		return
	}

	// 查询列表数据
	if err = query.Order(order).Offset(offsetSize).Limit(req.PageSize).Find(&orderList).Error; err != nil {
		return
	}

	// 计算总页数
	var totalPage int64 = 0
	if total > 0 {
		totalPage = int64(math.Ceil(float64(total) / float64(req.PageSize)))
	}

	orderPage = schemas.OrderPageRes{
		Items:     orderList,
		PageSize:  req.PageSize,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	return
}
