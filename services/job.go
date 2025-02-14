package services

import (
	"app-api/boot"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"encoding/json"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

const filename string = "jobPublish"

type job struct{}

func NewJob() *job {
	return &job{}
}

// 发布任务-订单支付
func (s *job) NewOrderPaidTask(data *schemas.OrderPaidPayload) (taskInfo *asynq.TaskInfo, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		utils.ZapLog().Error(filename, "NewOrderPaidTask", zap.Error(err))
		return nil, err
	}

	task := asynq.NewTask(consts.TypeOrderPaid, payload)
	if taskInfo, err = boot.Asynq.Client.Enqueue(task); err != nil {
		utils.ZapLog().Error(filename, "NewOrderPaidTask error", zap.Error(err))
		return
	}

	return
}

// 发布任务-任务完成
func (s *job) NewTaskCompletedTask(data *schemas.TaskCompletedPayload) (taskInfo *asynq.TaskInfo, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		utils.ZapLog().Error(filename, "NewTaskCompletedTask", zap.Error(err))
		return nil, err
	}

	task := asynq.NewTask(consts.TypeTaskCompleted, payload)
	if taskInfo, err = boot.Asynq.Client.Enqueue(task); err != nil {
		utils.ZapLog().Error(filename, "NewTaskCompletedTask error", zap.Error(err))
		return
	}

	return
}

// 发布任务-用户注册
func (s *job) NewUserRegisterTask(data *schemas.TaskUserRegisterPayload) (taskInfo *asynq.TaskInfo, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		utils.ZapLog().Error(filename, "NewUserRegisterTask", zap.Error(err))
		return nil, err
	}

	task := asynq.NewTask(consts.TypeUserRegister, payload)
	if taskInfo, err = boot.Asynq.Client.Enqueue(task); err != nil {
		utils.ZapLog().Error(filename, "NewUserRegisterTask error", zap.Error(err))
		return
	}

	return
}

// 发布任务-工具使用
func (s *job) NewToolUsedTask(data *schemas.TaskToolUsedPayload) (taskInfo *asynq.TaskInfo, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		utils.ZapLog().Error(filename, "NewToolUsedTask", zap.Error(err))
		return nil, err
	}

	task := asynq.NewTask(consts.TypeToolUsed, payload)
	if taskInfo, err = boot.Asynq.Client.Enqueue(task); err != nil {
		utils.ZapLog().Error(filename, "NewToolUsedTask error", zap.Error(err))
		return
	}

	return
}
