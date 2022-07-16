package routes

import (
	"antoccino/controllers"
	"antoccino/store"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, repo store.Store) {
	router.POST("/users", controllers.CreateUser(repo))
}
