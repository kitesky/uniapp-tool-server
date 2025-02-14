package jobs

import (
	"app-api/services"
	"app-api/types/schemas"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

// 工具使用类
func HandleToolUsedTask(ctx context.Context, t *asynq.Task) error {
	var p schemas.TaskToolUsedPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if err := services.NewTool().HandleToolUsedTask(&p); err != nil {
		return err
	}

	return nil
}
