package helpers

import (
	"antoccino/responses"
	"encoding/json"
	"log"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	var responseType string
	var innerData map[string]interface{}

	switch data.(type) {
	case error:
		log.Printf("an error occurred with statusCode %d: %v", statusCode, data.(error).Error())
		responseType = "error"
		innerData = map[string]interface{}{"error": data.(error).Error()}
	default:
		responseType = "success"
		innerData = map[string]interface{}{"data": data}
	}

	finalResponse := responses.UserResponse{
		Status:  statusCode,
		Message: responseType,
		Data:    innerData,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res, _ := json.Marshal(finalResponse)
	w.Write(res)
}
