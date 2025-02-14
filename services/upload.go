package services

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/models"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type upload struct{}

func NewUpload() *upload {
	return &upload{}
}

// 初始化文件路径
func (s *upload) initFilePath(req *schemas.UploadInitFilePathReq) (res *schemas.UploadInitFilePathRes, err error) {
	osDir, _ := os.Getwd()
	uuid := uuid.New().String()
	ext := ".png"
	filename := utils.GenerateStringUniqueID() + ext
	if req.Filename != "" {
		filename = req.Filename + ext
	}

	dir := schemas.DefaultDir
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

	return &schemas.UploadInitFilePathRes{
		Filename: filename,
		UserID:   req.UserID,
		Path:     path,
		FullPath: fullPath,
		UUID:     uuid,
		EXT:      ext,
	}, nil
}

func (s *upload) DownloadWechatAvatar(userID int64, avatarUrl string) (err error) {
	resp, err := http.Head(avatarUrl)
	if err != nil {
		err = errors.New("图片不存在")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		err = errors.New("图片不存在")
		return
	}
	response, err := http.Get(avatarUrl)
	if err != nil {
		err = errors.New("下载头像错误")
		return
	}
	defer response.Body.Close()

	customDir := "avatar"
	filepathInfo, err := s.initFilePath(&schemas.UploadInitFilePathReq{
		Dir:      customDir,
		Filename: utils.MD5([]byte(strconv.Itoa(int(userID)))),
		UserID:   userID,
	})

	file, err := os.Create(filepathInfo.FullPath)
	if err != nil {
		err = errors.New("创建文件失败")
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		err = errors.New("拷贝文件失败")
		return
	}

	// 更新用户头像
	if _, err = dao.NewUser().SetProfile(&schemas.UserProfileReq{
		UserID: userID,
		Avatar: boot.Config.App.AssetUrl + "/" + filepathInfo.Path,
	}); err != nil {
		err = errors.New("保存头像失败")
		return
	}

	return
}

func (s *upload) Upload(req schemas.UploadReq, ctx *gin.Context) (result schemas.UploadRes, err error) {
	file := req.File
	// 限制文件大小为5MB
	if file.Size > 5*1024*1024 {
		err = errors.New("上传文件最大限制5MB")
		return
	}

	// 获取文件类型
	mediaType, _, err := mime.ParseMediaType(file.Header.Get("Content-Type"))
	if err != nil {
		return
	}

	filename := filepath.Base(file.Filename)
	uuid := uuid.New().String()
	ext := filepath.Ext(filename)
	newFilename := utils.GenerateStringUniqueID() + ext
	if req.Filename != "" {
		newFilename = req.Filename + ext
	}
	dir := schemas.UploadAvatarDir
	if req.Dir != "" {
		dir = dir + req.Dir + "/"
	}

	path := dir + newFilename

	// 移动文件
	if err = ctx.SaveUploadedFile(file, path); err != nil {
		return
	}

	// 计算hash
	hash, _ := utils.MD5File(path)
	cfg := boot.Config.App
	fileModel, err := dao.NewFile().CreateFile(&models.File{
		Name:      filename,
		UserID:    req.UserID,
		Disk:      "go-local",
		Path:      path,
		URL:       cfg.AssetUrl + "/" + path,
		Size:      file.Size,
		Hash:      hash,
		UUID:      uuid,
		Extension: strings.TrimLeft(ext, "."),
		Mime:      mediaType,
		ExpiredAt: nil,
	})

	if err != nil {
		return
	}

	copier.Copy(&result, fileModel)
	return
}
