package dao

import (
	"app-api/boot"
	"app-api/models"
	"strconv"
	"time"
)

type checkIn struct{}

func NewCheckIn() *checkIn {
	return &checkIn{}
}

func (s *checkIn) CheckIn(userID int64) (result *models.UserCheckIn, err error) {
	nowTime := time.Now()
	today, _ := strconv.Atoi(nowTime.Format("20060102"))
	yesterday, _ := strconv.Atoi(nowTime.AddDate(0, 0, -1).Format("20060102"))
	var count int32 = 0

	data := models.UserCheckIn{}
	boot.DB.Where("user_id = ? AND date = ?", userID, today).First(&data)

	// 已经签到
	if data.ID > 0 {
		return &data, nil
	}

	// 查询昨天签到数据
	yesterdayData := models.UserCheckIn{}
	boot.DB.Where("user_id = ? AND date = ?", userID, yesterday).First(&yesterdayData)
	if yesterdayData.ID > 0 {
		count = yesterdayData.Count
	}

	// 签到日志
	result = &models.UserCheckIn{
		UserID: userID,
		Date:   int32(today),
		Count:  count + 1,
	}

	boot.DB.Create(result)
	return
}
