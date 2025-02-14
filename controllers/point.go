package controllers

import (
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type point struct{}

func NewPoint() *point {
	return &point{}
}

func (s *point) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/point", s.PointList)
}

func (s *point) PointList(c *gin.Context) {
	req := schemas.PointPageReq{
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

	pointList, err := services.NewPoint().GetPointList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(pointList).Success()
}
