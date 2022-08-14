package route

import (
	"antoccino/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users", controller.CreateUser())
}
