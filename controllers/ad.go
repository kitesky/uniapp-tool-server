package controllers

import (
	"app-api/pkg/response"
	"app-api/services"

	"github.com/gin-gonic/gin"
)

type ad struct{}

func NewAD() *ad {
	return &ad{}
}

func (s *ad) Router(router *gin.RouterGroup) {
	router.GET("/ad/:key", s.Info)
}

func (s *ad) Info(c *gin.Context) {
	key := c.Param("key")
	resp, _ := services.NewAD().GetAdSpace(key)
	if resp.ID == 0 {
		response.New(c).SetCode(404).SetMessage("资源未发现").Error()
		return
	}

	response.New(c).SetData(resp).Success()
}
