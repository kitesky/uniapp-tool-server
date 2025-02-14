package jobs

import (
	"app-api/services"
	"app-api/types/schemas"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

// 任务完成后处理
func HandleUserRegisterTask(ctx context.Context, t *asynq.Task) error {
	var p schemas.TaskUserRegisterPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if err := services.NewUser().HandleUserRegisterTask(&p); err != nil {
		return err
	}

	return nil
}
