package schemas

// 二维码初始化目录后返回
type QRCodePathRes struct {
	Filename string `json:"filename"`
	UserID   int64  `json:"user_id"`
	Path     string `json:"path"`
	FullPath string `json:"fullpath"`
	UUID     string `json:"uuid"`
	EXT      string `json:"ext"`
}

type QRCodeReq struct {
	UserID   int64  `json:"user_id"`
	Text     string `json:"text"`
	Filename string `json:"filename"`
	Dir      string `json:"dir"`
}
