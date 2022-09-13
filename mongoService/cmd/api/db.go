package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func(app *Config) ConnectMongo() *Config {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	_ = cancel
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to Mongo Successful...")
	}
	
	return &Config{
		Db: client,
	}

}

func(app *Config) GetMongoCollection(client *mongo.Client,collName string) *mongo.Collection {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collName)
	
	return collection
}