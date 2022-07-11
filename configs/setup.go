package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() (*mongo.Client, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Successfully connected to MongoDB!")

	return client, cancel
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DBName()).Collection(collectionName)
	return collection
}
