package schemas

import "app-api/models"

type HomeRes struct {
	Taxonomy  []*models.Taxonomy `json:"taxonomy"`
	Recommend []*models.AppTool  `json:"recommend"`
	Swiper    []*models.AdItem   `json:"swiper"`
	Option    []*models.Option   `json:"option"`
}
