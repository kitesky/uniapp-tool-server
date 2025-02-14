package services

import (
	"app-api/boot"
	"app-api/models"
)

type option struct{}

func NewOption() *option {
	return &option{}
}

func (s *option) GetOptionList(keyList []string) (list []*models.Option, err error) {
	err = boot.DB.Where("`key` in ?", keyList).Find(&list).Error
	return list, err
}
