package controllers

import (
	"app-api/dao"
	"app-api/pkg/response"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type product struct{}

func NewProduct() *product {
	return &product{}
}

func (p *product) Router(router *gin.RouterGroup) {
	router.GET("/product", p.GetProductList)
}

func (p *product) GetProductList(c *gin.Context) {
	req := schemas.ProductListReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.New(c).SetMessage("参数错误").Error()
		return
	}

	result, err := dao.NewProduct().GetProductList(req.Type)
	if err != nil {
		response.New(c).SetMessage("获取订单列表失败").Error()
		return
	}

	response.New(c).SetData(result).Success()
}
