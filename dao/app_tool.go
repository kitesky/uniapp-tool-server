package dao

import (
	"app-api/boot"
	"app-api/models"
)

type appTool struct{}

func NewAppTool() *appTool {
	return &appTool{}
}

func (s *appTool) GetAppTool(code string) (result *models.AppTool, err error) {
	result = &models.AppTool{}
	err = boot.DB.Where("code = ?", code).First(result).Error
	return
}
