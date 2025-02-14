package tasks

import (
	"app-api/types/consts"
)

// 任务处理器
var TaskhandlerList = map[string]TaskHandler{
	consts.TaskDailySignIn:      NewTaskCheckIn(),
	consts.TaskUserRegister:     NewTaskUserRegister(),
	consts.TaskToolUsedTextOnce: NewTaskToolUsedTextOnce(),
	consts.TaskUserInvite:       NewTaskUserInvite(),
	consts.TaskVIPGiftDaily:     NewTaskVIPGiftDaily(),
}
