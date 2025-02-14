package services

import (
	"app-api/boot"
	"app-api/models"

	"gorm.io/gorm"
)

type taxonomy struct{}

func NewTaxonomy() *taxonomy {
	return &taxonomy{}
}

func (s *taxonomy) GetAppTaxonomyListForHome() (list []*models.Taxonomy, err error) {
	list = make([]*models.Taxonomy, 0)
	tx := boot.DB
	err = tx.Preload("Meta").Where("taxonomy", "mini-app").Order("`order` asc").Find(&list).Error
	for index, item := range list {
		toolList := []*models.AppTool{}
		tx.Where("taxonomy_id = ? and status = ?", item.ID, "Y").Order("`sort` asc").Limit(4).Find(&toolList)

		list[index].Items = toolList
	}
	return list, err
}

func (s *taxonomy) GetAppTaxonomyListWithTool() (list []*models.Taxonomy, err error) {
	err = boot.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Where("status", "Y").Order("sort ASC")
	}).Preload("Meta").Where("taxonomy", "mini-app").Order("`order` asc").Find(&list).Error
	return list, err
}

func (s *taxonomy) GetAppTaxonomyWithTool(slug string) (info *models.Taxonomy, err error) {
	err = boot.DB.Preload("Items").Where("slug = ?", slug).First(&info).Error
	return info, err
}
