package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/schemas"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type taskStat struct{}

func NewTaskStat() *taskStat {
	return &taskStat{}
}

// 更新统计数据
func (s *taskStat) UpdateTaskStat(userID int64, code string) (taskStat *models.TaskStat, err error) {
	boot.DB.Transaction(func(tx *gorm.DB) error {
		// 查询统计数据
		if taskStat, err = s.GetStat(userID, code); err != nil {
			return err
		}

		// 更新统计数据
		taskStat.TodayCount++
		taskStat.WeekCount++
		taskStat.MonthCount++
		taskStat.YearCount++
		taskStat.Count++

		if err = boot.DB.Save(taskStat).Error; err != nil {
			return err
		}

		return nil
	})

	return
}

// 重置统计数据
func (s *taskStat) ResetTaskStat(taskStat *models.TaskStat) (newTaskStat *models.TaskStat, err error) {
	date := s.GetTaskDateVal()
	if taskStat.ID == 0 {
		return
	}

	newTaskStat = taskStat
	isNeedUpdate := false

	// 今日
	if newTaskStat.TodayDate != date.Today {
		newTaskStat.TodayCount = 0
		newTaskStat.TodayDate = date.Today
		isNeedUpdate = true
	}

	// 本周
	if newTaskStat.WeekDate != date.Week {
		newTaskStat.WeekCount = 0
		newTaskStat.WeekDate = date.Week
		isNeedUpdate = true
	}

	// 本月
	if newTaskStat.MonthDate != date.Month {
		newTaskStat.MonthCount = 0
		newTaskStat.MonthDate = date.Month
		isNeedUpdate = true
	}

	// 本年
	if newTaskStat.YearDate != date.Year {
		newTaskStat.YearCount = 0
		newTaskStat.YearDate = date.Year
		isNeedUpdate = true
	}

	// 需要时更新操作
	if isNeedUpdate {
		err = boot.DB.Save(newTaskStat).Error
	}

	return
}

// 获取格式日期
func (s *taskStat) GetTaskDateVal() (data *schemas.TaskStatDate) {
	now := time.Now()

	// 今天
	dateString := now.Format("20060102")
	today, _ := strconv.Atoi(dateString)
	// 本周
	year, _ := now.ISOWeek()
	firstDayOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())
	// 计算当前时间是今年的第几周
	week := int32(now.Sub(firstDayOfYear).Hours()/(24*7)) + 1

	data = &schemas.TaskStatDate{
		Today: int32(today),
		Week:  week,
		Month: int32(now.Month()),
		Year:  int32(now.Year()),
	}

	return
}

// 获取任务统计列表
func (s *taskStat) GetStatList(userID int64) (newList []*models.TaskStat, err error) {
	list := []*models.TaskStat{}
	err = boot.DB.Where("user_id = ?", userID).Find(&list).Error
	for _, item := range list {
		newItem, _ := s.ResetTaskStat(item)
		newList = append(newList, newItem)
	}

	return newList, err
}

// 获取任务统计
func (s *taskStat) GetStat(userID int64, code string) (stat *models.TaskStat, err error) {
	stat = &models.TaskStat{}
	err = boot.DB.Where("user_id = ? and code = ?", userID, code).Find(&stat).Error
	if stat.ID == 0 {
		task, _ := NewTask().GetTask(code)
		date := s.GetTaskDateVal()
		stat = &models.TaskStat{
			UserID:     userID,
			TaskID:     task.ID,
			Code:       task.Code,
			Type:       task.Type,
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
		stat, _ = s.ResetTaskStat(stat)
	}

	return
}

// 获取任务进度列表
func (s *taskStat) GetTaskProgressList(userID int64) (list map[string]schemas.TaskProgressRes, err error) {
	statList, _ := s.GetStatList(userID)
	list = make(map[string]schemas.TaskProgressRes)
	for _, item := range statList {
		list[item.Code] = *s.GetTaskProgress(item)
		// fmt.Println(111, item)
	}

	return list, err
}

// 获取任务进度
func (s *taskStat) GetTaskProgress(taskStat *models.TaskStat) (res *schemas.TaskProgressRes) {
	var status string = "N"
	var remainderCount int32 = 0 // 剩余任务次数
	var completedCount int32 = 0 // 完成任务次数
	res = &schemas.TaskProgressRes{
		ID:             taskStat.ID,
		Code:           taskStat.Code,
		CompletedCount: completedCount,
		RemainderCount: remainderCount,
		Completed:      status,
	}

	switch taskStat.Type {
	case "once":
		completedCount = taskStat.Count
	case "daily":
		completedCount = taskStat.TodayCount
	case "weekly":
		completedCount = taskStat.WeekCount
	case "monthly":
		completedCount = taskStat.MonthCount
	case "none":
		completedCount = taskStat.Count
	}

	// 判断任务是否完成&计算剩余任务次数
	task, err := NewTask().GetTask(taskStat.Code)
	if err != nil {
		return
	}

	if completedCount >= task.RewardCount {
		res.Completed = "Y"
	} else {
		res.RemainderCount = taskStat.Count - completedCount
	}

	return
}
