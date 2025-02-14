package dao

import (
	"app-api/boot"
	"app-api/models"
)

type search struct{}

func NewSearch() *search {
	return &search{}
}
func (s *search) GetList() (result []*models.SearchLog, err error) {
	result = []*models.SearchLog{}
	err = boot.DB.Limit(6).Order("sort ASC,count Desc").Find(&result).Error
	return
}

func (s *search) SearchLog(req *models.SearchLog) (result *models.SearchLog) {
	tx := boot.DB
	tx.Where("title", req.Title).First(&result)
	if result.ID == 0 {
		result = req
		result.Count = 1
		tx.Create(result)
	} else {
		result.Count++
		tx.Save(result)
	}

	return
}
