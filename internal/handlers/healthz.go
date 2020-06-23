package handlers

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/model"
)

// Health returns status of micro service
func Health(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := app.HealthzService.Health()

		if err != nil {
			ResponseWriter(w, http.StatusServiceUnavailable, err.Error(), nil)
			return
		}

		ResponseWriter(w, http.StatusOK, "", model.NewHealthStatus(result))
	}
}
