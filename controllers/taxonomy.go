package controllers

import (
	"app-api/pkg/response"
	"app-api/services"

	"github.com/gin-gonic/gin"
)

type taxonomy struct{}

func NewTaxonomy() *taxonomy {
	return &taxonomy{}
}

func (s *taxonomy) Router(router *gin.RouterGroup) {
	router.GET("taxonomy", s.GetList)
	router.GET("taxonomy/:slug", s.GetInfo)
}

func (s *taxonomy) GetList(c *gin.Context) {
	res, _ := services.NewTaxonomy().GetAppTaxonomyListWithTool()
	response.New(c).SetData(res).Success()
}

func (s *taxonomy) GetInfo(c *gin.Context) {
	slug := c.Param("slug")
	taxonomy, _ := services.NewTaxonomy().GetAppTaxonomyWithTool(slug)
	response.New(c).SetData(taxonomy).Success()
}
