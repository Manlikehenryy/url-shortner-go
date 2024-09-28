package helpers

import (

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	// Check for proxy headers (commonly used in production)
	
	// ip := c.Request.Header.Get("X-Real-IP")
	// if ip == "" {
	// 	ip = c.Request.Header.Get("X-Forwarded-For")
	// }
	// if ip == "" {
	// 	ip = c.ClientIP()
	// }
	
	return c.ClientIP()
}