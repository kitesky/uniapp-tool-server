package services

import (
	"app-api/dao"
	"app-api/models"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type order struct{}

func NewOrder() *order {
	return &order{}
}

// 订单支付成功处理
func (o *order) HandleOrderPaidTask(payload *schemas.OrderPaidPayload) (order *models.Order, err error) {
	utils.ZapLog().Error("service-order", "OrderPaidProcess", zap.Any("request", payload))

	order, err = dao.NewOrder().GetOrder(payload.OrderNo)
	if err != nil {
		return
	}

	switch order.OrderType {
	case "recharge":
		return dao.NewOrder().RechargeOrderProcess(payload.OrderNo, payload.PayOrderNo, payload.PayTime)
	case "vip", "point", "product":
		return dao.NewOrder().ProductOrderProcess(payload.OrderNo, payload.PayOrderNo, payload.PayTime)
		// todo
	case "application":
		// todo
	}

	return nil, errors.New("订单类型错误")
}

func (o *order) GetOrderList(req schemas.OrderPageReq) (newOrderData *schemas.OrderNewPageRes, err error) {
	orderData, err := dao.NewOrder().GetOrderList(req)
	newOrderData = &schemas.OrderNewPageRes{}
	copier.Copy(newOrderData, orderData)
	for index, item := range newOrderData.Items {
		newOrderData.Items[index].OrderTypeText = schemas.OrderTypeOptions[item.OrderType].Title
		newOrderData.Items[index].PayTypeText = schemas.OrderPayTypeOptions[item.PayType]
		newOrderData.Items[index].StatusText = schemas.OrderStatusOptions[item.Status]
		newOrderData.Items[index].StatusColor = schemas.OrderStatusColorOptions[item.Status]
	}

	return
}
