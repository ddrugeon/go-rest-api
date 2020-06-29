package handlers

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/model"
)

// Health returns status of micro service
func Version(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWriter(w, http.StatusOK, "", model.NewVersion(app.Version.GitCommit, app.Version.BuildDate, app.Version.Version))
	}
}
