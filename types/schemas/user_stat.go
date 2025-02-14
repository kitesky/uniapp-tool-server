package schemas

// 统计结果结构
type UserStatRes struct {
	Code       string  `json:"code"`
	UserID     string  `json:"user_id"`
	TodayDate  string  `json:"today_date"`
	WeekDate   string  `json:"week_date"`
	MonthDate  string  `json:"month_date"`
	YearDate   float64 `json:"year_date"`
	TodayCount int32   `json:"today_count"`
	WeekCount  string  `json:"week_count"`
	MonthCount string  `json:"month_count"`
	YearCount  string  `json:"year_count"`
	Count      string  `json:"count"`
}

// 任务统计日期结构
type UserStatDate struct {
	Today int32 `json:"today"`
	Week  int32 `json:"week"`
	Month int32 `json:"month"`
	Year  int32 `json:"year"`
}
