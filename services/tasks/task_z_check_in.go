package tasks

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type taskCheckIn struct{}

func NewTaskCheckIn() *taskCheckIn {
	return &taskCheckIn{}
}

// 任务前置检查
func (s *taskCheckIn) TaskCheck(userID int64, code string) (ok bool, err error) {
	// 查询任务统计信息
	taskStat, err := dao.NewTaskStat().GetStat(userID, code)
	if err != nil {
		utils.ZapLog().Error("taskCheckIn", "查询任务统计信息错误", zap.Error(err))
		return false, err
	}

	// 查询进度 任务完成直接返回false
	process := dao.NewTaskStat().GetTaskProgress(taskStat)
	if process.Completed == "Y" {
		return false, errors.New("任务已完成")
	}

	return true, nil
}

// 任务奖励
func (s *taskCheckIn) TaskReward(userID int64, code string) (ok bool, err error) {
	ok = false
	// 查询任务信息
	task, _ := dao.NewTask().GetActiveTask(code)
	if task.ID == 0 {
		utils.ZapLog().Error("taskCheckIn", "任务下架或者已经禁用", zap.Error(err))
		return
	}

	if err = boot.DB.Transaction(func(tx *gorm.DB) error {
		// 更新任务统计
		_, err := dao.NewTaskStat().UpdateTaskStat(userID, code)
		if err != nil {
			return err
		}

		// 发放奖励并记录奖励日志
		if _, err := dao.NewTask().TaskReward(tx, &schemas.TaskRewardReq{
			UserID:       userID,
			RewardType:   task.RewardType,
			TaskCode:     code,
			RewardAmount: task.RewardAmount,
			Title:        task.Title,
			Description:  task.Description,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		utils.ZapLog().Error("taskCheckIn", "任务奖励错误", zap.Error(err))
		return
	}

	ok = true
	return
}
