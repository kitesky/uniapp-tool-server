package services

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/models"
	"app-api/types/schemas"
	"app-api/utils"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type qr_code struct{}

func NewQRCode() *qr_code {
	return &qr_code{}
}

// 保存文件记录
func (q *qr_code) saveRecord(req *schemas.QRCodePathRes) (result *models.File, err error) {
	// 计算hash
	hash, _ := utils.MD5File(req.FullPath)
	// 获取文件大小
	fileStat, err := os.Stat(req.FullPath)
	if err != nil {
		return nil, err
	}

	fileSize := fileStat.Size()
	// 获取文件MIME类型
	mimeType := mime.TypeByExtension(filepath.Ext(req.FullPath))
	cfg := boot.Config.App
	result, err = dao.NewFile().CreateFile(&models.File{
		Name:      req.Filename,
		UserID:    0,
		Disk:      "go-local",
		Path:      req.Path,
		URL:       cfg.AssetUrl + "/" + req.Path,
		Size:      fileSize,
		Hash:      hash,
		UUID:      req.UUID,
		Extension: strings.TrimLeft(req.EXT, "."),
		Mime:      mimeType,
		ExpiredAt: nil,
	})

	return
}

// 初始化目录
func (q *qr_code) initDirPath(req *schemas.QRCodeReq) (result *schemas.QRCodePathRes, err error) {
	osDir, _ := os.Getwd()
	uuid := uuid.New().String()
	ext := ".png"
	filename := utils.GenerateStringUniqueID() + ext
	if req.Filename != "" {
		filename = req.Filename + ext
	}

	dir := schemas.QRCodeGenerateDir
	if req.Dir != "" {
		dir = dir + req.Dir + "/"
	}

	path := strings.Join([]string{dir, filename}, "")
	// 目录不存在，创建目录
	fullDir := strings.Join([]string{osDir, dir}, "/")
	_, err = os.Stat(fullDir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(fullDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	fullPath := filepath.Join(fullDir, filename)

	return &schemas.QRCodePathRes{
		Filename: filename,
		UserID:   req.UserID,
		Path:     path,
		FullPath: fullPath,
		UUID:     uuid,
		EXT:      ext,
	}, nil
}

func (q *qr_code) SaveWechatMiniAppQRCode(req *schemas.QRCodeReq) (result *models.File, err error) {
	fileInfo, err := q.initDirPath(req)
	if err != nil {
		return
	}

	file, err := os.Create(fileInfo.FullPath)
	if err != nil {
		return
	}
	defer file.Close()

	// 获取小程序二维码 二进制数据
	body, err := NewWechatApp().GetMiniAppQRCode(req.Text)

	// 写入二进制数据到文件
	_, err = file.Write(body)
	if err != nil {
		return
	}

	result, err = q.saveRecord(fileInfo)
	return
}

func (q *qr_code) GenerateInviteQRCode(req *schemas.QRCodeReq) (result *models.File, err error) {
	qr, err := qrcode.NewWith(req.Text)
	if err != nil {
		return
	}

	fileInfo, err := q.initDirPath(req)
	if err != nil {
		return
	}

	w, err := standard.New(
		fileInfo.FullPath,
		standard.WithLogoImageFilePNG("./public/logo.png"),
		standard.WithLogoSizeMultiplier(2),
	)
	if err != nil {
		return nil, err
	}

	// save file
	if err = qr.Save(w); err != nil {
		return nil, err
	}

	result, err = q.saveRecord(fileInfo)
	return
}
