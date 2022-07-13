package helpers

import (
	"antoccino/responses"
	"github.com/gin-gonic/gin"
	"log"
)

func ReturnResponse(c *gin.Context, data any, statusCode int) {
	var finalResponse any

	switch data.(type) {
	case error:
		log.Printf("an error occurred with statusCode %d: %v", statusCode, data.(error).Error())
		finalResponse = responses.UserResponse{
			Status: "error",
			Error: gin.H{
				"code":    statusCode,
				"message": data.(error).Error(),
			},
		}
	default:
		finalResponse = responses.UserResponse{
			Status: "success",
			Data:   data,
		}
	}

	c.JSON(statusCode, finalResponse)
}
