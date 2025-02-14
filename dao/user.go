package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
	"app-api/utils"
	"time"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

func (u *user) UpdateInviteInfo(userID int64, inviteID int64, inviteQRCode string) (meta models.UserMeta, err error) {
	meta = models.UserMeta{}
	err = boot.DB.Where("id = ?", userID).First(&meta).Error
	if meta.ID == 0 {
		meta = models.UserMeta{
			UserID:       userID,
			InviteID:     inviteID,
			InviteQrcode: inviteQRCode,
		}
		boot.DB.Create(&meta)
	} else {
		meta.InviteID = inviteID
		meta.InviteQrcode = inviteQRCode
		err = boot.DB.Save(&meta).Error
	}

	return
}

func (u *user) IsVIPActive(userID int64) (ok bool) {
	userVIP := &models.UserVip{}
	boot.DB.Where("user_id = ?", userID).Find(userVIP)
	if userVIP.ID == 0 {
		return false
	}

	if time.Now().Before(userVIP.ExpireTime) {
		return true
	}

	return false
}
func (u *user) UpdateUserVIP(userID int64, month int) (userVIP *models.UserVip, err error) {
	userVIP = &models.UserVip{}
	boot.DB.Where("user_id = ?", userID).Find(userVIP)
	if userVIP.ID == 0 {
		userVIP = &models.UserVip{
			UserID:     userID,
			Active:     "Y",
			ExpireTime: time.Now().AddDate(0, month, 0),
		}
		boot.DB.Create(userVIP)
	} else {
		var expireTime time.Time
		// 检查会员是否过期
		if time.Now().After(userVIP.ExpireTime) {
			expireTime = time.Now().AddDate(0, month, 0)
		} else {
			expireTime = userVIP.ExpireTime.AddDate(0, month, 0)
		}

		userVIP.Active = "Y"
		userVIP.ExpireTime = expireTime
		boot.DB.Save(userVIP)
	}

	return
}

func (u *user) GetUserBase(userID int64) (user *models.User) {
	boot.DB.Where("id = ?", userID).First(&user)
	return
}

func (u *user) GetUser(userID int64) (user *models.User) {
	boot.DB.Where("id = ?", userID).Preload("VIP").Preload("Meta").First(&user)
	return
}

func (u *user) GetUserByMMID(MMID int64) (user *models.User) {
	boot.DB.Where("mmid = ?", MMID).First(&user)
	return
}

func (u *user) GetUserByMail(email string) (user *models.User) {
	boot.DB.Where("email = ?", email).First(&user)
	return
}

func (u *user) GetUserByUsername(username string) (user *models.User) {
	boot.DB.Where("username = ?", username).First(&user)
	return
}

func (u *user) UpdatePassword(userID int64, password string) (ok bool) {
	encryption, err := utils.GenerateFromPassword(password)
	if err != nil {
		return false
	}

	boot.DB.Where("id = ?", userID).Update("password", encryption)
	return true
}

func (u *user) GetUserOpenID(userID int64) (openID string, err error) {
	err = boot.DB.Model(&models.UserWechatApp{}).Where("user_id = ?", userID).Pluck("openid", &openID).Error
	return
}

// 设置个人资料
func (u *user) SetProfile(req *schemas.UserProfileReq) (ok bool, err error) {
	data := models.User{ID: req.UserID}
	ok = true

	if req.Avatar != "" {
		data.Avatar = req.Avatar
	}

	if req.Name != "" {
		data.Name = req.Name
	}

	if req.Gender != "" {
		data.Gender = req.Gender
	}

	if req.Website != "" {
		data.Website = req.Website
	}

	if req.Bio != "" {
		data.Bio = req.Bio
	}

	if err = boot.DB.Updates(&data).Error; err != nil {
		ok = false
	}

	return
}
