package main

import (
	"crypto/tls"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/manlikehenryy/url-shortener-go/configs"
	"github.com/manlikehenryy/url-shortener-go/database"
	"github.com/manlikehenryy/url-shortener-go/helpers"
)

var rdb *redis.Client

func nit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:      configs.Env.REDIS_ADDRESS,  // Redis address (without rediss://)
		Username:  configs.Env.REDIS_USERNAME, // Username for Redis instance
		Password:  configs.Env.REDIS_PASSWORD, // Password for Redis instance
		TLSConfig: &tls.Config{},              // Enables SSL/TLS
	})
}

func main() {
	


	// Initialize configuration
	configs.Init()
	helpers.Initialize()

	// Connect to the database
	database.Connect()

	// Retrieve the port from the config
	port := configs.Env.PORT
	if port == "" {
		port = "8080" // Set a default port if not specified
	}

	// Create a new Gin router
	app := gin.New()

	// Apply middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// Set up routes
	// routes.Setup(app)

	// Start the server
	err := app.Run(":" + port)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
