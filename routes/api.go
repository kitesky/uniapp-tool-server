package routes

import (
	"app-api/controllers"
	"app-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(app *gin.Engine) {
	// 公共路由
	n := app.Group("/api/notify")
	controllers.NewNotify().Router(n)

	// Header 鉴权
	r := app.Group("/api", middlewares.Header()) // , middlewares.Header

	// Register routes.
	controllers.NewTest().Router(r)
	controllers.NewUser().Router(r)
	controllers.NewHome().Router(r)
	controllers.NewTaxonomy().Router(r)
	controllers.NewTool().Router(r)
	controllers.NewAD().Router(r)
	controllers.NewTask().Router(r)
	controllers.NewCheckIn().Router(r)
	controllers.NewOrder().Router(r)
	controllers.NewScore().Router(r)
	controllers.NewBalance().Router(r)
	controllers.NewReward().Router(r)
	controllers.NewPoint().Router(r)
	controllers.NewUserTransfer().Router(r)
	controllers.NewUpload().Router(r)
	controllers.NewProduct().Router(r)
	controllers.NewActivity().Router(r)
	controllers.NewSearch().Router(r)
}
