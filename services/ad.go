package services

import (
	"app-api/boot"
	"app-api/models"
	"app-api/pkg/cache"
	"encoding/json"
)

type ad struct{}

func NewAD() *ad {
	return &ad{}
}

func (s *ad) GetAdSpace(key string) (adSpace *models.AdSpace, err error) {
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
