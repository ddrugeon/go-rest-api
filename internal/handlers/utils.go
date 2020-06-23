package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/model"
)

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := model.NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}
