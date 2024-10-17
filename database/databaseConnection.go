package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// init() runs automatically during package initialization
func init() {
	Client = Dbinstance()
}

func Dbinstance() *mongo.Client {

	if Client != nil {
		return Client // Return existing client if already initialized
	}

	mongoDb := "mongodb+srv://test:Thepreacher1@cluster0.sdg0iit.mongodb.net/"
	fmt.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return Client

}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return Collection
}
