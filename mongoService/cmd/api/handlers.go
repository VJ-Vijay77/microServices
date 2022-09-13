package main

import (
	"context"
	"log"
	"mongo/models"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Config) Mongo(c *gin.Context) {
	c.JSON(200, "This is Mongo Db")
}

func (app *Config) Insert(c *gin.Context) {
	collection := app.GetMongoCollection(app.Db,os.Getenv("COLLECTION_NAME"))
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		log.Println(err)
		c.JSON(500, "Binding Failed")
		return
	}
	newuser := models.Users{
		Id:       primitive.NewObjectID(),
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
		Age:      users.Age,
	}
	res, err := collection.InsertOne(context.Background(), newuser)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Insertion failed")
		return
	}

	c.JSON(200, gin.H{
		"status": "Insertion successful",
		"item":   res,
	})

}
