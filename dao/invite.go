package dao

import (
	"app-api/boot"
	"app-api/models"
)

type invite struct{}

func NewInvite() *invite {
	return &invite{}
}

func (s *invite) GetInviteList(userID int64) (result []*models.UserInviteLog, err error) {
	result = []*models.UserInviteLog{}
	err = boot.DB.Where("invite_user_id = ?", userID).Preload("User").Find(&result).Error
	return
}

func (s *invite) CreateInvite(data *models.UserInviteLog) (result *models.UserInviteLog, err error) {
	result = data
	err = boot.DB.Create(&result).Error
	return
}
