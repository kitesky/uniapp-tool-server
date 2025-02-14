package controllers

import (
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type balance struct{}

func NewBalance() *balance {
	return &balance{}
}

func (s *balance) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/balance", s.BalanceList)
}

func (s *balance) BalanceList(c *gin.Context) {
	req := schemas.BalancePageReq{
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

	balanceList, err := services.NewBalance().GetBalanceList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(balanceList).Success()
}
