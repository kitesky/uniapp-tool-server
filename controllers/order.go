package controllers

import (
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type order struct{}

func NewOrder() *order {
	return &order{}
}

func (s *order) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/order", s.OrderList)
	token.POST("/order/create", s.CreateOrder)
}

// 创建订单
func (s *order) CreateOrder(c *gin.Context) {
	req := schemas.OrderCreateReq{
		UserID:  c.GetInt64("user_id"),
		PayType: "wechat",
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	order, err := dao.NewOrder().CreateOrder(&req)
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	openID, err := dao.NewUser().GetUserOpenID(req.UserID)
	if err != nil {
		response.New(c).SetMessage("用户不存在").Error()
		return
	}

	// 发送支付请求
	result, err := services.NewPayment().WechatMiniApp(&schemas.PaymentWechatMiniAppReq{
		Amount:      order.PayAmount,
		Description: order.Description,
		OpenID:      openID,
		OutTradeNo:  order.OrderNo,
	})

	if err != nil {
		response.New(c).SetMessage("获取预支付字符串失败").Error()
		return
	}

	response.New(c).SetData(result).Success()
}

func (s *order) OrderList(c *gin.Context) {
	req := schemas.OrderPageReq{
		UserID:    c.GetInt64("user_id"),
		Page:      consts.Page,
		PageSize:  consts.PageSize,
		SortField: consts.SortField,
		SortType:  consts.SortType,
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	orderList, err := services.NewOrder().GetOrderList(req)
	if err != nil {
		response.New(c).SetMessage("获取订单列表失败").Error()
		return
	}

	response.New(c).SetData(orderList).Success()
}
