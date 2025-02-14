package controllers

import (
	"app-api/services"
	"app-api/types/schemas"
	"app-api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type test struct{}

func NewTest() *test {
	return &test{}
}

func (s *test) Router(r *gin.RouterGroup) {
	r.GET("/test", s.Index)
}

func (s *test) Index(c *gin.Context) {
	// result, err := dao.NewUser().UpdateUserVIP(295, 1)
	// result, err := services.NewQRCode().GenerateInviteQRCode(&schemas.QRCodeReq{
	// 	Text:     "https://www.baidu.com",
	// 	UserID:   295,
	// 	Filename: "10001295",
	// 	Dir:      "invite",
	// })
	// result, err := services.NewJob().NewToolUsedTask(&schemas.TaskToolUsedPayload{
	// 	UserID: 295,
	// 	Code:   "tool:nzzj",
	// })
	// result, err := services.NewWechatApp().GetMiniAppQRCode("295")
	// fmt.Println(result, err, boot.Config.WechatPay)
	err := services.NewPayment().WechatTransfer(&schemas.PaymentWechatTransferReq{
		OutBillNo:       utils.GenerateStringUniqueID(),
		TransferSceneID: "1000",
		OpenID:          "o0Fme5KVq3lduxx9TL_iKME5J-qA",
		TransferAmount:  0.3,
		TransferRemark:  "测试转账",
	})
	fmt.Println(err)

	// response.New(c).SetMessage("hello world").SetData(result).Success()
}
