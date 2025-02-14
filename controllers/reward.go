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

type reward struct{}

func NewReward() *reward {
	return &reward{}
}

func (s *reward) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/reward", s.RewardList).Use(middlewares.Token())
	token.POST("/reward/exchange", s.RewardExchange).Use(middlewares.Token())
}

// 奖励金兑换现金&点数
func (s *reward) RewardExchange(c *gin.Context) {
	req := schemas.RewardExchangeReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	switch req.CashType {
	case "cash":
		if err := dao.NewReward().RewardExchange(&req); err != nil {
			response.New(c).SetMessage(err.Error()).Error()
			return
		}
	case "point":
		if err := dao.NewReward().RewardExchangePoint(&req); err != nil {
			response.New(c).SetMessage(err.Error()).Error()
			return
		}
	}

	response.New(c).SetMessage("提交成功").Success()
}

func (s *reward) RewardList(c *gin.Context) {
	req := schemas.RewardPageReq{
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

	rewardList, err := services.NewReward().GetRewardList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(rewardList).Success()
}
