package schemas

type CheckInRes struct {
	Count     int      `json:"count"`
	Date      int      `json:"date"`
	CreatedAt Datetime `json:"created_at"`
}
