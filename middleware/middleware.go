package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manlikehenryy/url-shortener-go/database"
	"github.com/manlikehenryy/url-shortener-go/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

func RateLimit(c *gin.Context) {
	ip := helpers.GetClientIP(c)
	log.Println("ipAddress: "+ip)

	limit := 2 // 10 clicks per minute
	key := fmt.Sprintf("rate_limit:%s", ip)

	// Increment request count
	count, err := database.RDB.Incr(ctx, key).Result()
	if err != nil {
		c.Next()
	}

	// Set expiration for the key if it's new
	if count == 1 {
		err := database.RDB.Expire(ctx, key, time.Minute).Err()
		if err != nil {
			c.Next()
		}
	}

	// If the count exceeds the limit, reject the request
	if count > int64(limit) {
		helpers.SendError(c, http.StatusTooManyRequests, "Rate limit exceeded")
		c.Abort() // Abort the request pipeline if it fails
		return
	}

	c.Next()
}

func IsAuthenticated(c *gin.Context) {
	// Retrieve the JWT token from the cookie
	cookie, err := c.Cookie("jwt")
	if err != nil {
		helpers.SendError(c, http.StatusUnauthorized, "Unauthorized: No JWT token provided")
		c.Abort() // Abort the request pipeline if authentication fails
		return
	}

	// Parse the JWT token from the cookie
	userIdStr, err := helpers.ParseJwt(cookie)
	if err != nil {
		helpers.SendError(c, http.StatusUnauthorized, "Unauthorized: Invalid JWT token")
		c.Abort() // Abort the request pipeline if token parsing fails
		return
	}

	// Convert the user ID string to primitive.ObjectID
	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		helpers.SendError(c, http.StatusUnauthorized, "Unauthorized: Invalid user ID")
		c.Abort() // Abort the request pipeline if user ID conversion fails
		return
	}

	// Store the user ID in the request context
	c.Set("userId", userId)

	// Continue to the next middleware or handler
	c.Next()
}
