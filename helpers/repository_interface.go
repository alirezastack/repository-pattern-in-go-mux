package helpers

import (
	"antoccino/configs"
	"antoccino/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository interface {
	CreateUser(ctx context.Context, u models.User) (string, error)
}

type UserMongo struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	Title    string             `bson:"title"`
}

func (c UserMongo) ToEntity() (models.User, error) {
	return models.User{
		Id:       c.Id.String(),
		Name:     c.Name,
		Location: c.Location,
		Title:    c.Title,
	}, nil
}

type MongoDBRepository struct {
	Client *mongo.Client
}

func (r *MongoDBRepository) CreateUser(ctx context.Context, user models.User) (string, error) {
	var usersCollection = configs.GetCollection(r.Client, "users")
	newUser := UserMongo{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	res, err := usersCollection.InsertOne(ctx, newUser)
	log.Printf("A new user is created with ID %s successfully", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
