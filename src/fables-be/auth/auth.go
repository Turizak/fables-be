package auth

import (
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, authorized := utilities.AuthorizeRequest(c)
		if !authorized {
			c.Abort() // Stop further processing
			return
		}

		// Store claims in context for use in handlers
		c.Set("claims", claims)

		// Continue to the next handler
		c.Next()
	}
}
