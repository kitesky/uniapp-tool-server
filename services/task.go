package services

import (
	"app-api/services/tasks"
	"app-api/utils"

	"go.uber.org/zap"
)

type task struct{}

func NewTask() *task {
	return &task{}
}

// 任务处理
func (s *task) HandleTaskCompletedTask(userID int64, code string) (ok bool, err error) {
	// 获取任务操作句柄
	handler, err := tasks.NewTaskHandler().GetTaskHandle(code)
	if err != nil {
		utils.ZapLog().Info("task", "获取任务句柄错误", zap.Error(err))
		return
	}

	// 检查任务是否完成
	if ok, err = handler.TaskCheck(userID, code); !ok {
		utils.ZapLog().Info("task", "任务检查不通过", zap.Error(err))
		return
	}

	// 发放奖励
	if ok, err = handler.TaskReward(userID, code); !ok {
		utils.ZapLog().Info("task", "发放奖励失败", zap.Error(err))
		return
	}

	return true, nil
}
