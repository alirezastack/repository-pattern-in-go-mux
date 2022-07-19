package store

import (
	"antoccino/configs"
	"antoccino/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoStore struct {
	store *mongo.Client
}

func (mc *mongoStore) CreateUser(user models.User) (string, error) {
	var usersCollection = GetCollection(mc.store, "users")
	newUser := models.UserMongo{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// release resources if CreateUser completes before timeout elapses
	defer cancel()

	res, err := usersCollection.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}

	log.Printf("A new user is created with ID %s successfully", res.InsertedID)

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// NewMongoDBStore returns a MongoDB store
func NewMongoDBStore() Store {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.MongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Successfully connected to MongoDB!")

	s := &mongoStore{
		store: client,
	}
	return s
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(configs.DBName()).Collection(collectionName)
	return collection
}
