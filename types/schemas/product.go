package schemas

type ProductListReq struct {
	Type string `json:"type" form:"type"`
}

type ProductVipExtra struct {
	Month int   `json:"month"`
	Point int64 `json:"point"` // 赠送点数
	Score int64 `json:"score"` // 赠送积分
}

type ProductPointExtra struct {
	Point int64 `json:"point"` // 点数
	Gift  int64 `json:"gift"`  // 赠送点数
}
