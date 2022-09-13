package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// var MONGO_URI = "mongodb://localhost:27017"
//
//	var MONGO_URI = "mongodb://mongodb:27017"
var PORT = ":8081"

type Config struct {
	Db *mongo.Client
}

func main() {
	var app = &Config{}
	
	
	app = app.ConnectMongo()

	r := gin.Default()
	app.Routess(r)
	log.Println("Starting MongoDB Service on Port ", PORT)

	err := r.Run(PORT)
	if err != nil {
		log.Panic(err)
	}
}
