package dao

import (
	"app-api/boot"
	"app-api/models"
	"app-api/types/consts"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type task struct{}

func NewTask() *task {
	return &task{}
}

// 获取任务列表
func (s *task) GetTaskList(userID int64) (list []*schemas.TaskListRes, err error) {
	taskList := []*models.Task{}
	err = boot.DB.Where("status = ? and app_status = ?", "Y", "Y").Find(&taskList).Error
	if err != nil {
		return
	}

	// 计算任务进度
	copier.Copy(&list, &taskList)
	taskProgressList, _ := NewTaskStat().GetTaskProgressList(userID)
	for index, item := range list {
		stat := taskProgressList[item.Code]
		if stat.ID > 0 {
			list[index].Completed = stat.Completed
			list[index].CompletedCount = stat.CompletedCount
			list[index].RemainderCount = stat.RemainderCount
		} else {
			list[index].Completed = "N"
			list[index].CompletedCount = 0
			list[index].RemainderCount = item.RewardCount
		}
	}

	return list, err
}

// 获取任务信息
func (s *task) GetTask(code string) (task *models.Task, err error) {
	task = &models.Task{}
	err = boot.DB.Where("code = ?", code).First(task).Error
	return
}

// 获取可用任务信息
func (s *task) GetActiveTask(code string) (task *models.Task, err error) {
	task = &models.Task{}
	err = boot.DB.Where("code = ? and status = ?", code, "Y").First(task).Error
	return
}

// 任务发放奖励
func (s *task) TaskReward(tx *gorm.DB, req *schemas.TaskRewardReq) (ok bool, err error) {
	switch req.RewardType {
	case consts.TaskRewardTypePoint:
		pointLog := &models.UserPointLog{
			UserID:      req.UserID,
			Type:        consts.ScoreActionInc,
			Code:        req.TaskCode,
			Amount:      int64(req.RewardAmount),
			Title:       req.Title,
			Description: req.Description,
		}

		if _, err := NewPoint().PointLog(tx, pointLog); err != nil {
			return false, err
		}
	case consts.TaskRewardTypeScore:
		scoreLog := &models.UserScoreLog{
			UserID:      req.UserID,
			Type:        consts.ScoreActionInc,
			Code:        req.TaskCode,
			Amount:      int64(req.RewardAmount),
			Title:       req.Title,
			Description: req.Description,
		}

		if _, err := NewScore().ScoreLog(tx, scoreLog); err != nil {
			return false, err
		}
	case consts.TaskRewardTypeBalance:
		balanceLog := &models.UserBalanceLog{
			UserID:      req.UserID,
			Type:        consts.BalanceActionInc,
			Code:        req.TaskCode,
			Amount:      req.RewardAmount,
			Title:       req.Title,
			Description: req.Description,
		}

		if _, err := NewBalance().BalanceLog(tx, balanceLog); err != nil {
			return false, err
		}
	case consts.TaskRewardTypeReward:
		rewardLog := &models.UserRewardLog{
			UserID:      req.UserID,
			Type:        consts.RewardActionInc,
			Code:        req.TaskCode,
			Amount:      req.RewardAmount,
			Title:       req.Title,
			Description: req.Description,
		}

		if _, err := NewReward().RewardLog(tx, rewardLog); err != nil {
			return false, err
		}
	}

	return true, nil
}
