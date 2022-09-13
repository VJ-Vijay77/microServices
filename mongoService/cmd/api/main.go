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
	// client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Connection to Mongo Successful...")
	// }
	
	app = ConnectMongo()

	r := gin.Default()
	app.Routess(r)
	log.Println("Starting MongoDB Service on Port ", PORT)

	err := r.Run(PORT)
	if err != nil {
		log.Panic(err)
	}
}
