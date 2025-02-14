package tasks

import "errors"

type taskHandler struct{}

func NewTaskHandler() *taskHandler {
	return &taskHandler{}
}

func (h *taskHandler) GetTaskHandle(code string) (handler TaskHandler, err error) {
	taskCode := string(code)
	if handler, ok := TaskhandlerList[taskCode]; ok {
		return handler, nil
	} else {
		err = errors.New("任务配置不正确")
	}

	return
}
