package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manlikehenryy/url-shortener-go/helpers"
	"github.com/manlikehenryy/url-shortener-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUrl(c *gin.Context) {
	var url models.Url

	if err := c.ShouldBindJSON(&url); err != nil {
		log.Println("Unable to parse body:", err)
		helpers.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	userId, ok := c.MustGet("userId").(primitive.ObjectID)
	if !ok {
		helpers.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	url.UserId = userId

	insertResult, err := urlCollection.InsertOne(context.Background(), url)
	if err != nil {
		log.Println("Database error:", err)
		helpers.SendError(c, http.StatusInternalServerError, "Failed to create url")
		return
	}

	url.ID = insertResult.InsertedID.(primitive.ObjectID)

	helpers.SendJSON(c, http.StatusCreated, gin.H{
		"data":    url,
		"message": "Url created successfully",
	})
}

func GetAllUrl(c *gin.Context) {
	var urls []models.Url
	filter := bson.M{}

	params, err := helpers.PaginateCollection(c, urlCollection, filter, &urls)
	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve urls")
		return
	}

	helpers.SendPaginatedResponse(c, urls, params)
}

func GetUrl(c *gin.Context) {

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helpers.SendError(c, http.StatusBadRequest, "Invalid url ID")
		return
	}

	userId, ok := c.MustGet("userId").(primitive.ObjectID)
	if !ok {
		helpers.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	var url models.Url
	err = urlCollection.FindOne(context.Background(), bson.M{"_id": id, "userId": userId}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.SendError(c, http.StatusNotFound, "Url not found")
		} else {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve url")
		}
		return
	}

	helpers.SendJSON(c, http.StatusOK, gin.H{
		"data": url,
	})
}

func UpdateUrl(c *gin.Context) {

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helpers.SendError(c, http.StatusBadRequest, "Invalid url ID")
		return
	}

	userId, ok := c.MustGet("userId").(primitive.ObjectID)
	if !ok {
		helpers.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	var url models.Url
	if err := c.ShouldBindJSON(&url); err != nil {
		helpers.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	var existingUrl models.Url
	err = urlCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&existingUrl)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.SendError(c, http.StatusNotFound, "Url not found")
		} else {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve url")
		}
		return
	}
	

	if existingUrl.UserId != userId {
		helpers.SendError(c, http.StatusForbidden, "Unauthorized to update this url")
		return
	}

	update := bson.M{
		"$set": bson.M{
			"originalUrl":  url.OriginalUrl,
			"expiration":   url.Expiration,
		},
	}

	result, err := urlCollection.UpdateOne(context.Background(), bson.M{"_id": id, "userId": userId}, update)
	if err != nil {
		log.Println("Database error:", err)
		helpers.SendError(c, http.StatusInternalServerError, "Failed to update url")
		return
	}

	if result.MatchedCount == 0 {
		helpers.SendError(c, http.StatusNotFound, "Url not found or unauthorized")
		return
	}

	helpers.SendJSON(c, http.StatusOK, gin.H{
		"message": "Url updated successfully",
	})
}

func UsersUrl(c *gin.Context) {
	userId, ok := c.MustGet("userId").(primitive.ObjectID)
	if !ok {
		helpers.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	filter := bson.M{"userId": userId}

	
	var urls []models.Url
	params, err := helpers.PaginateCollection(c, urlCollection, filter, &urls)
	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve urls")
		return
	}

	helpers.SendPaginatedResponse(c, urls, params)
}

func DeleteUrl(c *gin.Context) {

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helpers.SendError(c, http.StatusBadRequest, "Invalid url ID")
		return
	}

	userId, ok := c.MustGet("userId").(primitive.ObjectID)
	if !ok {
		helpers.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	var existingUrl models.Url
	err = urlCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&existingUrl)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.SendError(c, http.StatusNotFound, "Url not found")
		} else {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve url")
		}
		return
	}

	if existingUrl.UserId != userId {
		helpers.SendError(c, http.StatusForbidden, "Unauthorized to delete this url")
		return
	}

	filter := bson.M{"_id": id, "userId": userId}
	result, err := urlCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to delete url")
		return
	}

	if result.DeletedCount == 0 {
		helpers.SendError(c, http.StatusNotFound, "Url not found or unauthorized")
		return
	}

	helpers.SendJSON(c, http.StatusOK, gin.H{
		"message": "Url deleted successfully",
	})
}
