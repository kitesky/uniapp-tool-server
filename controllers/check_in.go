package controllers

import (
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type checkIn struct{}

func NewCheckIn() *checkIn {
	return &checkIn{}
}

func (s *checkIn) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.POST("/check-in", s.CheckIn)
}

func (s *checkIn) CheckIn(c *gin.Context) {
	userID := c.GetInt64("user_id")

	data, err := dao.NewCheckIn().CheckIn(userID)
	if err != nil {
		response.New(c).SetMessage("签到失败").Error()
		return
	}

	checkInRes := schemas.CheckInRes{}
	copier.Copy(&checkInRes, &data)

	// 完成签到任务
	services.NewJob().NewTaskCompletedTask(&schemas.TaskCompletedPayload{
		TaskCode: consts.TaskDailySignIn,
		UserID:   userID,
	})

	response.New(c).SetData(checkInRes).Success()
}
