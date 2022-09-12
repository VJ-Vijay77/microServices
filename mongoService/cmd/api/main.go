package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var MONGO_URI = "mongodb://localhost:27017"
// var MONGO_URI = "mongodb://mongodb:27017"
var PORT = ":8081"
type Config struct{}
func main() {
	app := Config{}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx,cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(database)

	srv := &http.Server{
		Addr: PORT,
		Handler: app.Routes(),
	}

	log.Println("Starting MongoDB Service on Port ",PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
