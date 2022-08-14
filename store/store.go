package store

import (
	"antoccino/model"
)

// UserStore is a data storage interface for user
type UserStore interface {
	// CreateUser returns the ID of newly created user
	CreateUser(user model.User) (string, error)
}
