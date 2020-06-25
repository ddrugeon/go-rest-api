package handlers

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/model"
	"github.com/gorilla/mux"
)

func GetAllDroids(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, _ := app.DroidService.Get(r.Context())

		data := make([]model.Droid, 0)
		for _, current := range result {
			data = append(data, model.NewDroid(current.ID, current.Name, current.Type, current.Company, current.Class, current.Model, current.Height, current.Vehicles))
		}

		ResponseWriter(w, http.StatusOK, "", data)
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

		data := model.NewDroid(current.ID, current.Name, current.Type, current.Company, current.Class, current.Model, current.Height, current.Vehicles)
		ResponseWriter(w, http.StatusOK, "", data)
	}
}
