package controllers

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/schemas"
	"app-api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

func (s *user) Router(router *gin.RouterGroup) {
	r := router.Group("/user")
	r.POST("/sign-up", s.SignUp)                                  // 注册
	r.POST("/sign-in", s.SignIn)                                  // 登录
	r.POST("/sign-in-with-wechat-miniapp", s.WechatMiniAppSignIn) // 微信登录
	
	token := r.Use(middlewares.Token())
	token.POST("/change-password", s.ChangePassword)
	token.GET("/info", s.Info)
	token.POST("/profile", s.Profile)
	token.GET("/invite", s.GetMyInvite)
	token.GET("/payment-account/:type", s.GetPaymentAccount)    // 收款信息
	token.GET("/payment-account-list", s.GetPaymentAccountList) // 收款信息列表
	token.POST("/set-payment-account", s.SetPaymentAccount)     // 设置收款信息

}

// 微信小程序登录
func (s *user) WechatMiniAppSignIn(c *gin.Context) {
	req := &schemas.WechatAppLoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Println(11111, req)
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	utils.ZapLog().Info("user", "controller.WechatMiniAppSignIn 接收请求:", zap.Any("req", req))

	// 获取微信小程序的openid
	wechatResp, _ := services.NewWechatApp().Code2Session(req.Code)
	if wechatResp.Openid == "" {
		response.New(c).SetMessage("兑换微信唯一身份ID失败").Error()
		return
	}

	req.OpenID = wechatResp.Openid
	resultRes, err := services.NewUser().WechatAppSignIn(req)
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).SetData(resultRes).Success()
}

func (s *user) GetMyInvite(c *gin.Context) {
	userID := c.GetInt64("user_id")
	resp, err := services.NewInvite().GetMyInvite(userID)
	if err == nil {
		response.New(c).SetData(&resp).Success()
		return
	}

	response.New(c).SetMessage(err.Error()).Error()
}

func (s *user) Profile(c *gin.Context) {
	req := schemas.UserProfileReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	_, err := dao.NewUser().SetProfile(&req)
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).Success()
}

func (s *user) Info(c *gin.Context) {
	req := schemas.UserInfoReq{UserID: c.GetInt64("user_id")}
	userRes := schemas.UserInfoRes{
		VIP: &schemas.UserInfoVIPRes{Active: "N"},
	}

	// 会员权益
	userRes.VIP.Benefits = []*schemas.UserInfoVIPBenefits{}
	for _, benefit := range schemas.UserVIPBenefits {
		newBenefit := *benefit
		newBenefit.Icon = boot.Config.App.AssetUrl + benefit.Icon
		userRes.VIP.Benefits = append(userRes.VIP.Benefits, &newBenefit)
	}

	if user, _ := services.NewUser().GetUser(req.UserID); user.ID > 0 {
		copier.Copy(&userRes, &user)
		response.New(c).SetData(userRes).Success()
		return
	}

	response.New(c).SetMessage("用户不存在").SetCode(401).Error()
}

func (s *user) SignUp(c *gin.Context) {
	req := schemas.SignUpReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	if err := services.NewUser().SignUp(&req); err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).Success()
}

func (s *user) SignIn(c *gin.Context) {
	req := schemas.SignInReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	result, err := services.NewUser().SignIn(&req)
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).SetData(result).Success()
}

func (s *user) ChangePassword(c *gin.Context) {
	req := schemas.ChangePasswordReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	if err := services.NewUser().ChangePassword(&req); err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).SetMessage("修改成功").Success()
}

func (s *user) SetPaymentAccount(c *gin.Context) {
	req := schemas.UserPaymentAccountReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	result := dao.NewUserPaymentAccount().SetPaymentAccount(&req)
	response.New(c).SetData(result).Success()
}

func (s *user) GetPaymentAccountList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	result := dao.NewUserPaymentAccount().GetPaymentAccountList(userID)
	response.New(c).SetData(result).Success()
}

func (s *user) GetPaymentAccount(c *gin.Context) {
	payType := c.Param("type")
	userID := c.GetInt64("user_id")
	result := dao.NewUserPaymentAccount().GetPaymentAccount(userID, payType)
	response.New(c).SetData(result).Success()
}
