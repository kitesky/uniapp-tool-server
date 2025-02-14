package dao

import (
	"app-api/boot"
	"app-api/models"
)

type file struct{}

func NewFile() *file {
	return &file{}
}

func (s *file) CreateFile(file *models.File) (result *models.File, err error) {
	err = boot.DB.Create(&file).Error
	return file, err
}
