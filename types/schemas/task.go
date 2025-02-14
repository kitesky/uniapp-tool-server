package schemas

type TaskHandlerReq struct {
	UserID int64  `json:"-" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

// 任务列表结构
type TaskListRes struct {
	Code           string  `json:"code"`
	Title          string  `json:"title"`
	LaunchTitle    string  `json:"launch_title"`
	LaunchURL      string  `json:"launch_url"`
	LaunchType     string  `json:"launch_type"`
	CompletedTitle string  `json:"completed_title"`
	RewardType     string  `json:"reward_type"`
	RewardAmount   float64 `json:"reward_amount"`
	RewardCount    int32   `json:"reward_count"`
	RewardIcon     string  `json:"reward_icon"`
	Icon           string  `json:"icon"`
	Description    string  `json:"description"`
	Content        string  `json:"content"`
	Payload        string  `json:"payload"`
	RemainderCount int32   `json:"remainder_count"`
	CompletedCount int32   `json:"completed_count"`
	Completed      string  `json:"completed"`
}

// 任务进度结构
type TaskProgressRes struct {
	ID             int64  `json:"id"`
	Code           string `json:"code"`
	RemainderCount int32  `json:"remainder_count"`
	CompletedCount int32  `json:"completed_count"`
	Completed      string `json:"completed"`
}

// 任务统计日期结构
type TaskStatDate struct {
	Today int32 `json:"today"`
	Week  int32 `json:"week"`
	Month int32 `json:"month"`
	Year  int32 `json:"year"`
}

type TaskRewardReq struct {
	UserID       int64   `json:"user_id"`
	RewardType   string  `json:"reward_type"`
	TaskCode     string  `json:"task_code"`
	RewardAmount float64 `json:"reward_amount"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
}
