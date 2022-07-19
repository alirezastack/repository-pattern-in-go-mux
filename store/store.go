package store

import (
	"antoccino/models"
)

// Store is a data storage interface
type Store interface {
	// CreateUser returns the ID of newly created user
	CreateUser(user models.User) (string, error)
}
