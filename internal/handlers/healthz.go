package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
)

// Health returns status of micro service
func Health(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		response, _ := json.Marshal(app.HealthzService.Health())
		w.Write(response)
	}
}
