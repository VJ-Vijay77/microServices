package main

import (
	"log"

	"github.com/gin-gonic/gin"
)


func(app *Config) Mongo(c *gin.Context) {
	c.JSON(200,"This is Mongo Db")
}

func(app *Config) Insert(c *gin.Context) {
	collection := app.GetMongoCollection()
	log.Println(collection.Name())
	
}