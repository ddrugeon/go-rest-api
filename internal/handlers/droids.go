package handlers

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/gorilla/mux"
)

func GetAllDroids(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, _ := app.DroidService.Get(r.Context())

		ResponseWriter(w, http.StatusOK, "", result)
	}
}

func GetDroidByID(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params = mux.Vars(r)
		id := params["id"]

		current, err := app.DroidService.GetByID(r.Context(), id)
		if err != nil {
			ResponseWriter(w, http.StatusNotFound, err.Error(), nil)
			return
		}

		ResponseWriter(w, http.StatusOK, "", current)
	}
}
