package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
}

type UserMongo struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	Title    string             `bson:"title"`
}

func (c UserMongo) ToMongoEntity() (User, error) {
	return User{
		Id:       c.Id.Hex(),
		Name:     c.Name,
		Location: c.Location,
		Title:    c.Title,
	}, nil
}
