package database

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/manlikehenryy/url-shortener-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

var RDB *redis.Client

// Initialize Redis

func init() {
	if configs.Env.MODE == "production" {
		RDB = redis.NewClient(&redis.Options{
			Addr: configs.Env.REDIS_ADDRESS, // Redis address (without rediss://)
		})
	} else {
		RDB = redis.NewClient(&redis.Options{
			Addr:      configs.Env.REDIS_ADDRESS,  // Redis address (without rediss://)
			Username:  configs.Env.REDIS_USERNAME, // Username for Redis instance
			Password:  configs.Env.REDIS_PASSWORD, // Password for Redis instance
			TLSConfig: &tls.Config{},              // Enables SSL/TLS
		})
	}

}

// Initialize the MongoDB client
func Connect() {
	mongoURI := configs.Env.MONGO_DB_URI
	if mongoURI == "" {
		log.Fatal("MONGO_DB_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = Client.Database("go_url_shortener")

	log.Println("Connected to MongoDB")
}
