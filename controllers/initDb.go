package controllers

import "go.mongodb.org/mongo-driver/mongo"

var usersCollection *mongo.Collection
var urlCollection *mongo.Collection


func InitDB(DB *mongo.Database) {

	usersCollection = DB.Collection("users")
	urlCollection = DB.Collection("url")
}
