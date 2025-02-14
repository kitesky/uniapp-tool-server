package controllers

import (
	"app-api/pkg/response"
	"app-api/services"
	"app-api/types/schemas"

	"github.com/gin-gonic/gin"
)

type home struct{}

func NewHome() *home {
	return &home{}
}

func (s *home) Router(router *gin.RouterGroup) {
	router.GET("/home", s.Info)
}

func (s *home) Info(c *gin.Context) {
	res := schemas.HomeRes{}
	res.Taxonomy, _ = services.NewTaxonomy().GetAppTaxonomyListForHome()
	res.Recommend, _ = services.NewTool().GetRecommentTool()

	swiper, _ := services.NewAD().GetAdSpace("mini-app-home-swiper")
	res.Swiper = swiper.Items
	// res.Option, _ = services.NewOption().GetOptionList([]string{"meta_title", "meta_keywords", "meta_description"})
	response.New(c).SetData(res).Success()
}
