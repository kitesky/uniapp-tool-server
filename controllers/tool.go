package controllers

import (
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/consts"
	"app-api/types/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tool struct{}

func NewTool() *tool {
	return &tool{}
}

func (s *tool) Router(router *gin.RouterGroup) {
	router.GET("tool", s.GetToolList)
	router.GET("tool/:id", s.GetTool)

	token := router.Group("/").Use(middlewares.Token())
	token.POST("tool/handler", s.Handler) // 发送处理请求 需要登录
}

func (s *tool) Handler(c *gin.Context) {
	req := schemas.ToolHandlerReq{UserID: c.GetInt64("user_id")}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	// 处理请求
	result, err := services.NewTool().ToolHandler(req.UserID, req.Code, req.Data)
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	// 推送到队列中
	services.NewJob().NewToolUsedTask(&schemas.TaskToolUsedPayload{
		UserID: req.UserID,
		Code:   req.Code,
	})

	response.New(c).SetData(result).Success()
}

func (s *tool) GetToolList(c *gin.Context) {
	req := schemas.ToolPageReq{
		Page:      consts.Page,
		PageSize:  consts.PageSize,
		SortField: consts.SortField,
		SortType:  consts.SortType,
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	toolList, err := services.NewTool().GetToolList(req)
	if err != nil {
		response.New(c).SetMessage("获取列表失败").Error()
		return
	}

	response.New(c).SetData(toolList).Success()
}

func (s *tool) GetTool(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tool, err := services.NewTool().GetTool(int64(id))
	if err != nil {
		response.New(c).SetMessage("获取详情失败").Error()
		return
	}

	response.New(c).SetData(tool).Success()
}
