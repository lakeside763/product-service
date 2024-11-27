package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, err error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var errorMessage string
	if err != nil {
		errorMessage = fmt.Sprintf(message, err)
	} else {
		errorMessage = message
	}

	response := map[string]string{
		"error": errorMessage,
	}

	json.NewEncoder(w).Encode(response)
}