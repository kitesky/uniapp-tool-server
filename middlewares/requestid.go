package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// New initializes the RequestID middleware.
func RequestID() gin.HandlerFunc {
	headerXRequestID := "X-Request-ID"
	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(headerXRequestID)
		if rid == "" {
			rid = uuid.New().String()
			c.Request.Header.Add(headerXRequestID, rid)
		}

		// Set the id to ensure that the requestid is in the response
		c.Header(headerXRequestID, rid)
		c.Next()
	}
}
