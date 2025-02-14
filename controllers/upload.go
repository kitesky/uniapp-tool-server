package controllers

import (
	"app-api/dao"
	"app-api/middlewares"
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/schemas"
	"app-api/utils"

	"github.com/gin-gonic/gin"
)

type upload struct{}

func NewUpload() *upload {
	return &upload{}
}

func (s *upload) Router(router *gin.RouterGroup) {
	token := router.Group("upload").Use(middlewares.Token())
	token.POST("/avatar", s.UploadAvatar)
}

func (s *upload) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	userID := c.GetInt64("user_id")
	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	// 上传文件
	result, err := services.NewUpload().Upload(schemas.UploadReq{
		ExpiredAt: nil,
		File:      file,
		UserID:    userID,
		Filename:  utils.MD5([]byte(c.GetString("user_id"))),
	}, c)

	if err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	// 更新用户头像
	if _, err = dao.NewUser().SetProfile(&schemas.UserProfileReq{
		UserID: userID,
		Avatar: result.URL,
	}); err != nil {
		response.New(c).SetMessage(err.Error()).Error()
		return
	}

	response.New(c).SetData(result).Success()
}
