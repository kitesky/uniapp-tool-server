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

type score struct{}

func NewScore() *score {
	return &score{}
}

func (s *score) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/score", s.ScoreList).Use(middlewares.Token())
	token.POST("/score/exchange", s.ScoreExchange).Use(middlewares.Token())
}

// 积分兑换现金
func (s *score) ScoreExchange(c *gin.Context) {
	req := schemas.ScoreExchangeReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	if err := dao.NewScore().ScoreExchange(&req); err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).SetMessage("兑换成功").Success()
}

func (s *score) ScoreList(c *gin.Context) {
	req := schemas.ScorePageReq{
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

	scoreList, err := services.NewScore().GetScoreList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(scoreList).Success()
}
