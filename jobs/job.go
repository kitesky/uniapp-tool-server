package jobs

import (
	"app-api/types/consts"

	"github.com/hibiken/asynq"
)

func JobSchedule(mux *asynq.ServeMux) {
	mux.HandleFunc(TypeWelcomeEmail, HandleWelcomeEmailTask)
	mux.HandleFunc(TypeReminderEmail, HandleReminderEmailTask)

	// 已支付订单任务
	mux.HandleFunc(consts.TypeOrderPaid, HandleOrderPaidTask)
	// 任务完成后触发
	mux.HandleFunc(consts.TypeTaskCompleted, HandleTaskCompletedTask)
	// 用户注册
	mux.HandleFunc(consts.TypeUserRegister, HandleUserRegisterTask)
	// 工具使用
	mux.HandleFunc(consts.TypeToolUsed, HandleToolUsedTask)
}
