package store

import (
	"antoccino/config"
	"antoccino/model"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoUserStore struct {
	store *mongo.Client
}

func (mc *mongoUserStore) CreateUser(user model.User) (string, error) {
	var usersCollection = GetCollection(mc.store, "users")
	newUser := model.UserMongo{
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

	log.Info().Msgf("A new user is created with ID %s successfully", res.InsertedID)

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// NewMongoDBStore returns a MongoDB store
func NewMongoDBStore() UserStore {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.MongoURI()))
	if err != nil {
		event := log.Fatal()
		event.Msg(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		event := log.Fatal()
		event.Msg(err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		event := log.Fatal()
		event.Msg(err.Error())
	}

	log.Info().Msg("Successfully connected to MongoDB!")

	s := &mongoUserStore{
		store: client,
	}
	return s
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(configs.DBName()).Collection(collectionName)
	return collection
}
