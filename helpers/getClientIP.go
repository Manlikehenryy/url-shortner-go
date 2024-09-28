package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	// Check for proxy headers (commonly used in production)
	log.Println("X-Real-IP: "+c.Request.Header.Get("X-Real-IP"))
	log.Println("X-Forwarded-For: "+c.Request.Header.Get("X-Forwarded-For"))
	log.Println("ClientIP: "+c.ClientIP())
	
	ip := c.Request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.ClientIP()
	}
	return ip
}