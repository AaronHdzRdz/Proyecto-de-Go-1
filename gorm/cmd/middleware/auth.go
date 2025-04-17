package middleware

import "github.com/gin-gonic/gin"

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("x-api-key")
		if apikey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API key is required"})
			return
		}
		if apikey != "ashdasdhasjhdgajshdg" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key"})
			return
		}
		c.Next()
	}
}
