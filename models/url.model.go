package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ShortUrl     string             `json:"shortUrl" bson:"shortUrl"`
	OriginalUrl  string             `json:"originalUrl" bson:"originalUrl" binding:"required"`
	Expiration   int64              `json:"expiration" bson:"expiration" binding:"required"` //in seconds
	ClickCount   int                `bson:"clickCount"`
	ClickDetails []Click            `bson:"clickDetails"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
}

type Click struct {
	IPAddress string    `bson:"ipAddress"`
	Timestamp time.Time `bson:"timestamp"`
}
