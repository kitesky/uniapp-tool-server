package controllers

import (
	"app-api/services"
	"app-api/types/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type notify struct{}

func NewNotify() *notify {
	return &notify{}
}

func (s *notify) Router(router *gin.RouterGroup) {
	router.POST("/wechat-miniapp-pay", s.WechatMiniAppPay)
}

// 微信小程序支付回调
func (s *notify) WechatMiniAppPay(c *gin.Context) {
	if _, err := services.NewPayment().WechatMiniAppNotify(c.Request); err != nil {
		c.JSON(http.StatusOK, schemas.PaymentWechatPayCallbackRes{
			Code:    "FAIL",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schemas.PaymentWechatPayCallbackRes{
		Code:    "SUCCESS",
		Message: "成功",
	})
}
