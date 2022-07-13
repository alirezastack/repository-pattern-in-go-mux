package routes

import (
	"antoccino/controllers"
	"antoccino/helpers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, repo *helpers.MongoDBRepository) {
	router.POST("/users", controllers.CreateUser(repo))
}
