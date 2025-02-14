package consts

// 队列名称
const (
	OrderQueue = "critical"
	TaskQueue  = "critical"
)

// 任务类型列表
const (
	// order
	TypeOrderCreated = "order:created" // 订单创建
	TypeOrderPaid    = "order:paid"    // 订单支付

	// task
	TypeTaskCompleted = "task:completed" // 任务完成

	// user
	TypeUserRegister = "user:register" // 新用户注册

	// tool
	TypeToolUsed = "tool:used" // 工具使用
)
