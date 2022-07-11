package controllers

import (
	"antoccino/helpers"
	"antoccino/models"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

var validate = validator.New()

func CreateUser(repo helpers.UserRepository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var user models.User

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			helpers.ReturnResponse(rw, err, http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&user); err != nil {
			helpers.ReturnResponse(rw, err, http.StatusBadRequest)
			return
		}

		newUser := models.User{
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}

		insertedId, err := repo.CreateUser(ctx, newUser)
		if err != nil {
			helpers.ReturnResponse(rw, err, http.StatusInternalServerError)
			return
		}
		log.Printf("A new user is created with ID %s successfully", insertedId)

		helpers.ReturnResponse(rw, map[string]interface{}{"id": insertedId}, http.StatusCreated)
	}
}
