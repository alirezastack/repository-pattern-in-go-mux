package store

import (
	"antoccino/models"
	"context"
)

// Store is a data storage interface
type Store interface {
	// CreateUser returns the ID of newly created user
	CreateUser(ctx context.Context, user models.User) (string, error)
}
