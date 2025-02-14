package jobs

import (
	"app-api/services"
	"app-api/types/schemas"
	"app-api/utils"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// 处理订单支付
func HandleOrderPaidTask(ctx context.Context, t *asynq.Task) error {
	var p schemas.OrderPaidPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		utils.ZapLog().Error("order_queue", "json.Unmarshal 错误", zap.Error(err))
		return err
	}

	if _, err := services.NewOrder().HandleOrderPaidTask(&p); err != nil {
		utils.ZapLog().Error("order_queue", "OrderPaidProcess 错误", zap.Error(err))
		return err
	}

	return nil
}
