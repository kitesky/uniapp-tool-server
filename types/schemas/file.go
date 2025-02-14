package schemas

import (
	"mime/multipart"
	"time"
)

const (
	DefaultDir        = "storege/files/"
	UploadDir         = "storege/files/uploads/"
	UploadAvatarDir   = "storege/files/avatar/"
	QRCodeGenerateDir = "storege/files/qrcode/"
)

type UploadInitFilePathReq struct {
	UserID   int64  `json:"user_id"`
	Filename string `json:"filename"`
	Dir      string `json:"dir"`
}

type UploadInitFilePathRes struct {
	UserID   int64  `json:"user_id"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
	FullPath string `json:"fullpath"`
	UUID     string `json:"uuid"`
	EXT      string `json:"ext"`
}

type UploadReq struct {
	File      *multipart.FileHeader `json:"file" binding:"required"`
	UserID    int64                 `json:"user_id"`
	Filename  string                `json:"filename"`
	Dir       string                `json:"dir"`
	ExpiredAt *time.Time            `json:"expired_at"`
}

type UploadRes struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
	URL  string `json:"url"`
}
