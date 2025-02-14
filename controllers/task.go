package controllers

import (
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type task struct{}

func NewTask() *task {
	return &task{}
}

func (s *task) Router(router *gin.RouterGroup) {
	token := router.Group("/").Use(middlewares.Token())
	token.GET("/task", s.GetTaskList)
	token.POST("/task/handler", s.TaskHandler)
}

// 领取任务奖励
func (s *task) TaskHandler(c *gin.Context) {
	req := schemas.TaskHandlerReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误[code]").Error()
		return
	}

	ok, err := services.NewTask().HandleTaskCompletedTask(req.UserID, req.Code)
	if ok {
		response.New(c).SetMessage("领取成功").Success()
		return
	}

	response.New(c).SetMessage(err.Error()).Error()
}

func (s *task) GetTaskList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	resp, _ := dao.NewTask().GetTaskList(userID)
	if len(resp) == 0 {
		response.New(c).SetCode(404).SetMessage("暂无任务").Error()
		return
	}

	response.New(c).SetData(resp).Success()
}
