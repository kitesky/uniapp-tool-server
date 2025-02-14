package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
)

type userPaymentAccount struct{}

func NewUserPaymentAccount() *userPaymentAccount {
	return &userPaymentAccount{}
}

func (u *userPaymentAccount) GetPaymentAccountList(userID int64) (result []*models.UserPaymentAccount) {
	boot.DB.Where("user_id = ?", userID).Find(&result)
	return
}

func (u *userPaymentAccount) GetPaymentAccount(userID int64, payType string) (result *models.UserPaymentAccount) {
	boot.DB.Where("user_id = ? and pay_type = ?", userID, payType).First(&result)
	return
}

func (u *userPaymentAccount) SetPaymentAccount(req *schemas.UserPaymentAccountReq) (result *models.UserPaymentAccount) {
	result = &models.UserPaymentAccount{
		UserID:  req.UserID,
		PayType: req.PayType,
		Account: req.Account,
		Name:    req.Name,
	}

	paymentInfo := &models.UserPaymentAccount{}
	boot.DB.Where("user_id = ? and pay_type = ?", req.UserID, req.PayType).First(&paymentInfo)

	if paymentInfo.ID == 0 {
		boot.DB.Create(&result)
	} else {
		boot.DB.Model(&paymentInfo).Updates(&result)
	}

	return
}
