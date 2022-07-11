package routes

import (
	"antoccino/controllers"
	"antoccino/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func UserRoute(router *mux.Router, repo helpers.UserRepository) {
	router.HandleFunc("/users", controllers.CreateUser(repo)).Methods(http.MethodPost)
}
