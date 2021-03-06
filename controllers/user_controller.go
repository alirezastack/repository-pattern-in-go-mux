package controllers

import (
	"antoccino/helpers"
	"antoccino/models"
	"antoccino/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var validate = validator.New()

func CreateUser(repo store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBind(&user); err != nil {
			helpers.ReturnResponse(c, err, http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&user); err != nil {
			helpers.ReturnResponse(c, err, http.StatusBadRequest)
			return
		}

		log.Printf("createUser payload: %+v", user)

		newUser := models.User{
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}

		insertedId, err := repo.CreateUser(newUser)
		if err != nil {
			helpers.ReturnResponse(c, err, http.StatusInternalServerError)
			return
		}

		helpers.ReturnResponse(c, gin.H{"id": insertedId}, http.StatusCreated)
	}
}
