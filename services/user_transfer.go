package services

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/models"
	"app-api/pkg/cache"
	"app-api/types/schemas"
	"encoding/json"

	"github.com/jinzhu/copier"
)

type userTransfer struct{}

func NewUserTransfer() *userTransfer {
	return &userTransfer{}
}

func (s *userTransfer) GetAdSpace(key string) (adSpace *models.AdSpace, err error) {
	var cacheKeys = cache.GetAdKeys(key)
	resp, err := cache.Get(cacheKeys.Key)

	if err != nil {
		if err = boot.DB.Where("`key`", key).Preload("Items").Find(&adSpace).Error; err == nil {
			jsonStr, _ := json.Marshal(adSpace)
			cache.Set(cacheKeys.Key, string(jsonStr), cacheKeys.TTL)
			return
		}

		return
	}

	// 解析数据
	json.Unmarshal([]byte(resp), &adSpace)
	return
}

func (s *userTransfer) GetUserTransferList(req schemas.UserTransferPageReq) (newResult *schemas.UserTransferNewPageRes, err error) {
	result, err := dao.NewUserTransfer().GetUserTransferList(req)
	newResult = &schemas.UserTransferNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].PayTypeText = schemas.UserTransferPayTypeOptions[item.PayType]
		newResult.Items[index].StatusText = schemas.UserTransferStatusOptions[item.Status].Text
		newResult.Items[index].StatusColor = schemas.UserTransferStatusOptions[item.Status].Color
	}

	return
}
