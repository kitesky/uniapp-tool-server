package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		// c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Max-Age", "172800")
		// c.Header("Access-Control-Allow-Credentials", "false")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok")
		}

		c.Next()
	}
}
