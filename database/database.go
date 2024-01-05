package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"golang-review-phone/config"
)

var collection *mongo.Collection

func InitDB() {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("golang_backend_review")
	collection = database.Collection("reviews")

	log.Println("Connected to MongoDB!")

	// Defer a call to Disconnect after instantiating your client
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func GetCollection(collectionName string) *mongo.Collection {
	return collection.Database().Collection(collectionName)
}
