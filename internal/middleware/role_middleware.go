package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")
		if !exist || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Admin access only",
			})

			return
		}
		c.Next()
	}
}
