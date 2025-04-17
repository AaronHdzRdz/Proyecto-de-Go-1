package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		//antes del handler
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		log.Printf("Request: %s %s from %s", method, path, clientIP)
		c.Next()
		//Despues del handler
		entTime:= time.Now()
		duration := entTime.Sub(startTime)
		statusCode := c.Writer.Status()
		log.Printf("Response: %d %s in %v", statusCode, path, duration)
	}
}