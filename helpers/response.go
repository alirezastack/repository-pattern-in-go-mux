package helpers

import (
	"antoccino/responses"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ReturnResponse(c *gin.Context, data any, statusCode int) {
	var finalResponse any

	switch data.(type) {
	case error:
		log.Error().Msgf("an error occurred with statusCode %d: %v", statusCode, data.(error).Error())
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

func GinCustomRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		log.Error().Msgf("An internal server error occurred: %s", err)
		ReturnResponse(c, errors.New(err), http.StatusInternalServerError)
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
