package controllers

import (
	"app-api/dao"
	"app-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type search struct{}

func NewSearch() *search {
	return &search{}
}

func (s *search) Router(router *gin.RouterGroup) {
	router.GET("/search", s.GetHotList)
}

func (s *search) GetHotList(c *gin.Context) {
	list, _ := dao.NewSearch().GetList()
	response.New(c).SetData(list).Success()
}
