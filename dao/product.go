package dao

import (
	"app-api/boot"
	"app-api/models"
)

type product struct{}

func NewProduct() *product {
	return &product{}
}

func (p *product) GetProduct(id int64) (product *models.Product, err error) {
	product = &models.Product{}
	err = boot.DB.Where("id = ?", id).First(product).Error
	return
}

func (p *product) GetProductList(productType string) (products []*models.Product, err error) {
	products = []*models.Product{}
	err = boot.DB.Where("`type` = ? and status = ?", productType, "Y").Find(&products).Error
	return
}
