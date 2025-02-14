package schemas

import "time"

// 订单支付数据
type OrderPaidPayload struct {
	PayOrderNo string    `json:"pay_order_no"`
	PayTime    time.Time `json:"pay_time"`
	OrderNo    string    `json:"order_no"`
	Amount     string    `json:"amount"`
}

// 任务完成数据
type TaskCompletedPayload struct {
	UserID   int64  `json:"user_id"`
	TaskCode string `json:"task_code"`
}

// 新用户注册
type TaskUserRegisterPayload struct {
	UserID   int64 `json:"user_id"`
	InviteID int64 `json:"invite_id"`
}

// 工具使用
type TaskToolUsedPayload struct {
	UserID int64  `json:"user_id"`
	Code   string `json:"code"`
}
