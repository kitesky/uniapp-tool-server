package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type userStat struct{}

func NewUserStat() *userStat {
	return &userStat{}
}

// 更新统计数据
func (s *userStat) UpdateUserStat(userID int64, code string) (userStat *models.UserStat, err error) {
	boot.DB.Transaction(func(tx *gorm.DB) error {
		// 查询统计数据
		if userStat, err = s.GetStat(userID, code); err != nil {
			return err
		}

		// 更新统计数据
		userStat.TodayCount++
		userStat.WeekCount++
		userStat.MonthCount++
		userStat.YearCount++
		userStat.Count++

		if err = boot.DB.Save(userStat).Error; err != nil {
			return err
		}

		return nil
	})

	return
}

// 重置统计数据
func (s *userStat) ResetUserStat(userStat *models.UserStat) (newUserStat *models.UserStat, err error) {
	date := s.GetUserStatDateVal()
	if userStat.ID == 0 {
		return
	}

	newUserStat = userStat
	isNeedUpdate := false

	// 今日
	if newUserStat.TodayDate != date.Today {
		newUserStat.TodayCount = 0
		newUserStat.TodayDate = date.Today
		isNeedUpdate = true
	}

	// 本周
	if newUserStat.WeekDate != date.Week {
		newUserStat.WeekCount = 0
		newUserStat.WeekDate = date.Week
		isNeedUpdate = true
	}

	// 本月
	if newUserStat.MonthDate != date.Month {
		newUserStat.MonthCount = 0
		newUserStat.MonthDate = date.Month
		isNeedUpdate = true
	}

	// 本年
	if newUserStat.YearDate != date.Year {
		newUserStat.YearCount = 0
		newUserStat.YearDate = date.Year
		isNeedUpdate = true
	}

	// 需要时更新操作
	if isNeedUpdate {
		err = boot.DB.Save(newUserStat).Error
	}

	return
}

// 获取格式日期
func (s *userStat) GetUserStatDateVal() (data *schemas.UserStatDate) {
	now := time.Now()

	// 今天
	dateString := now.Format("20060102")
	today, _ := strconv.Atoi(dateString)
	// 本周
	year, _ := now.ISOWeek()
	firstDayOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())
	// 计算当前时间是今年的第几周
	week := int32(now.Sub(firstDayOfYear).Hours()/(24*7)) + 1

	data = &schemas.UserStatDate{
		Today: int32(today),
		Week:  week,
		Month: int32(now.Month()),
		Year:  int32(now.Year()),
	}

	return
}

// 获取任务统计列表
func (s *userStat) GetStatList(userID int64) (newList []*models.UserStat, err error) {
	list := []*models.UserStat{}
	err = boot.DB.Where("user_id = ?", userID).Find(&list).Error
	for _, item := range list {
		newItem, _ := s.ResetUserStat(item)
		newList = append(newList, newItem)
	}

	return newList, err
}

// 获取任务统计
func (s *userStat) GetStat(userID int64, code string) (stat *models.UserStat, err error) {
	stat = &models.UserStat{}
	err = boot.DB.Where("user_id = ? and code = ?", userID, code).Find(&stat).Error
	if stat.ID == 0 {
		task, _ := NewTask().GetTask(code)
		date := s.GetUserStatDateVal()
		stat = &models.UserStat{
			UserID:     userID,
			Code:       task.Code,
			TodayDate:  date.Today,
			WeekDate:   date.Week,
			MonthDate:  date.Month,
			TodayCount: 0,
			WeekCount:  0,
			MonthCount: 0,
			Count:      0,
		}

		err = boot.DB.Create(stat).Error
	} else {
		// 验证任务统计日期，不一致的重置统计
		stat, _ = s.ResetUserStat(stat)
	}

	return
}
