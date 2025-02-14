package routes

import (
	"github.com/gin-gonic/gin"
)

// registers the routes and middlewares necessary for the server
func RegisterWebRoutes(app *gin.Engine) {
	app.GET("/h5", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
}
