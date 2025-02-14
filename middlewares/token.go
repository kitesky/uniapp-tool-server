package middlewares

import (
	"app-api/pkg/response"
	"app-api/utils"

	"github.com/gin-gonic/gin"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		// 参数不能空
		if token == "" {
			response.New(c).SetCode(401).SetMessage("未登录").Error()
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			response.New(c).SetCode(401).SetMessage(err.Error()).Error()
			c.Abort()
			return
		}

		// 设置userId
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
