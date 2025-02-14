package controllers

import (
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type userTransfer struct{}

func NewUserTransfer() *userTransfer {
	return &userTransfer{}
}

func (s *userTransfer) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/transfer", s.UserTransferList)
}

func (s *userTransfer) UserTransferList(c *gin.Context) {
	req := schemas.UserTransferPageReq{
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

	userTransferList, err := services.NewUserTransfer().GetUserTransferList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(userTransferList).Success()
}
