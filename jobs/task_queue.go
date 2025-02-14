package jobs

import (
	"app-api/services"
	"app-api/types/schemas"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

// 任务完成后处理
func HandleTaskCompletedTask(ctx context.Context, t *asynq.Task) error {
	var p schemas.TaskCompletedPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if _, err := services.NewTask().HandleTaskCompletedTask(p.UserID, p.TaskCode); err != nil {
		return err
	}

	return nil
}
