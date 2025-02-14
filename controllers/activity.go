package controllers

import (
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"

	"github.com/gin-gonic/gin"
)

type activity struct{}

func NewActivity() *activity {
	return &activity{}
}

func (s *activity) Router(router *gin.RouterGroup) {
	router.GET("/activity/:uuid", s.GetActivity)

	token := router.Group("/").Use(middlewares.Token())
	token.GET("/activity", s.ActivityList)
}

func (s *activity) GetActivity(c *gin.Context) {
	uuid := c.Param("uuid")
	activity, _ := dao.NewActivity().GetActivityWithUUID(uuid)
	if activity.ID == 0 {
		response.New(c).SetCode(404).SetMessage("资源未发现").Error()
		return
	}

	// 剔除markdown 返回纯文本
	activity.Content = utils.StripOptions(activity.Content)

	response.New(c).SetData(activity).Success()
}

func (s *activity) ActivityList(c *gin.Context) {
	req := schemas.ActivityPageReq{
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

	activityList, err := services.NewActivity().GetActivityList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(activityList).Success()
}
