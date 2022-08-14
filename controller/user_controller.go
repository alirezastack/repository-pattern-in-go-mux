package controller

import (
	"antoccino/helper"
	"antoccino/model"
	"antoccino/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

var validate = validator.New()
var userRepo store.UserStore

func init() {
	userRepo = store.NewMongoDBStore()
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.ShouldBind(&user); err != nil {
			helper.ReturnResponse(c, err, http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&user); err != nil {
			helper.ReturnResponse(c, err, http.StatusBadRequest)
			return
		}

		log.Info().Msgf("createUser payload: %+v", user)

		newUser := model.User{
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}

		insertedId, err := userRepo.CreateUser(newUser)
		if err != nil {
			helper.ReturnResponse(c, err, http.StatusInternalServerError)
			return
		}

		helper.ReturnResponse(c, gin.H{"id": insertedId}, http.StatusCreated)
	}
}
