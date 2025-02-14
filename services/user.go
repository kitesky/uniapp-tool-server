package services

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/models"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

// 用户注册队列任务
func (s *user) HandleUserRegisterTask(req *schemas.TaskUserRegisterPayload) error {
	// 绑定邀请关系
	if req.InviteID != 0 {
		user := dao.NewUser().GetUserByMMID(req.InviteID)
		// 绑定邀请关系
		dao.NewInvite().CreateInvite(&models.UserInviteLog{
			UserID:       req.UserID,
			InviteUserID: user.ID,
		})

		// 完成分享邀请积分奖励任务
		NewJob().NewTaskCompletedTask(&schemas.TaskCompletedPayload{
			TaskCode: consts.TaskUserInvite,
			UserID:   user.ID,
		})
	}

	// 完成新人福利 1元体验金任务
	NewJob().NewTaskCompletedTask(&schemas.TaskCompletedPayload{
		TaskCode: consts.TaskUserRegister,
		UserID:   req.UserID,
	})

	return nil
}

// 查询用户信息
func (s *user) GetUser(userID int64) (user *models.User, err error) {
	user = dao.NewUser().GetUser(userID)
	if user.ID == 0 {
		return
	}

	// MMID
	if user.MMID == 0 {
		user.MMID = decimal.NewFromInt(consts.UserMinIDStart + user.ID).IntPart()
		boot.DB.Save(&user)
	}

	// VIP
	if user.VIP.ID == 0 {
		user.VIP = models.UserVip{Active: "N"}
	} else {
		// 检查会员是否过期
		if time.Now().After(user.VIP.ExpireTime) {
			user.VIP.Active = "N"
			boot.DB.Model(&user.VIP).Where("user_id = ?", userID).Update("active", "N")
		}
	}

	// 邀请码
	if user.Meta.InviteID == 0 {
		inviteQrcode, _ := NewQRCode().SaveWechatMiniAppQRCode(&schemas.QRCodeReq{
			Text:     fmt.Sprintf("invite_id=%d", user.MMID),
			UserID:   userID,
			Filename: strconv.Itoa(int(user.MMID)),
			Dir:      "invite",
		})
		dao.NewUser().UpdateInviteInfo(userID, user.MMID, inviteQrcode.URL)
	}

	// 图像地址URL
	prefixURL := boot.Config.App.AssetUrl
	if user.Avatar == "" {
		user.Avatar = prefixURL + consts.UserDefaultAvatar
	}

	return
}

func (s *user) SignUp(req *schemas.SignUpReq) (err error) {
	encryption, err := utils.GenerateFromPassword(req.Password)
	if err != nil {
		return err
	}

	userInfo := dao.NewUser().GetUserByUsername(req.Username)
	if userInfo.ID > 0 {
		return errors.New("用户名已存在")
	}

	userInfo = dao.NewUser().GetUserByMail(req.Username)
	if userInfo.ID > 0 {
		return errors.New("邮箱已存在")
	}

	data := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: encryption,
	}

	return boot.DB.Create(&data).Error
}

func (s *user) SignIn(req *schemas.SignInReq) (res schemas.SignInRes, err error) {
	userInfo := dao.NewUser().GetUserByUsername(req.Username)
	if userInfo.ID == 0 {
		return res, errors.New("用户名不存在")
	}

	isVerify := utils.CompareHashAndPassword(userInfo.Password, req.Password)
	if !isVerify {
		return res, errors.New("密码错误")
	}

	result, err := utils.GenerateToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return res, err
	}

	res = schemas.SignInRes{
		AccessToken: result.AccessToken,
		ExpiresAt:   result.ExpiresAt,
		IssuedAt:    result.IssuedAt,
		ExpiresIn:   result.ExpiresIn,
	}

	return
}

func (s *user) ChangePassword(req *schemas.ChangePasswordReq) (err error) {
	if req.ConfirmPassword != req.Password {
		return errors.New("两次密码不一致")
	}

	userInfo := dao.NewUser().GetUser(req.UserID)
	if userInfo.ID == 0 {
		return errors.New("用户不存在")
	}

	isVerify := utils.CompareHashAndPassword(userInfo.Password, req.OldPassword)
	if !isVerify {
		return errors.New("原密码错误")
	}

	dao.NewUser().UpdatePassword(req.UserID, req.Password)
	return
}

// 微信小程序登录
func (s *user) WechatAppSignIn(req *schemas.WechatAppLoginReq) (res schemas.SignInRes, err error) {
	wechatApp := models.UserWechatApp{}
	user := models.User{}
	boot.DB.Where("openid = ?", req.OpenID).First(&wechatApp)
	if wechatApp.ID == 0 {
		// 新用户注册
		randomStr := utils.GenerateRandomString(6)
		password, _ := utils.GenerateFromPassword(randomStr)
		uniqID := utils.GenerateStringUniqueID()
		user = models.User{
			Name:     "微信用户_" + randomStr,
			Email:    "virtual." + uniqID + "@idcd.com",
			MMID:     0,
			Gender:   "female",
			Password: password,
		}
		boot.DB.Create(&user)
		boot.DB.Create(&models.UserWechatApp{UserID: user.ID, Openid: req.OpenID})

		// 推送到队列中
		NewJob().NewUserRegisterTask(&schemas.TaskUserRegisterPayload{
			UserID:   user.ID,
			InviteID: req.InviteID,
		})

		// 下载用户头像
		go NewUpload().DownloadWechatAvatar(user.ID, req.AvatarUrl)
	} else {
		// 已注册用户
		boot.DB.Where("id = ?", wechatApp.UserID).First(&user)
	}

	// 生成token
	result, err := utils.GenerateToken(user.ID, user.Name)
	if err == nil {
		res = schemas.SignInRes{
			AccessToken: result.AccessToken,
			ExpiresAt:   result.ExpiresAt,
			IssuedAt:    result.IssuedAt,
			ExpiresIn:   result.ExpiresIn,
		}
	}

	return
}
