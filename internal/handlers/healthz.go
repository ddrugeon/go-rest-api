package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
)

// HealthStatus represents current health of microservice and it is used to serialize response.
type healthStatus struct {
	Status string `json:"status"`
}

// Health returns status of micro service
func Health(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		result, err := app.HealthzService.Health()
		status := healthStatus{Status: ""}
		if err != nil {
			w.WriteHeader(503)
			status.Status = err.Error()

		} else {
			w.WriteHeader(200)
			status.Status = result
		}

		response, _ := json.Marshal(status)
		w.Write(response)
	}
}
