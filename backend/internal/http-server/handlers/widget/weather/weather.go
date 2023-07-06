package weather

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	Weather *entities.Weather `json:"weather"`
}

type Getter interface {
	GetWeather() (*entities.Weather, error)
}

func GetWeather(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO weather by region
		weather, err := getter.GetWeather()

		if err != nil {
			log.Error("failed to get weather", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get weather"))

			return
		}

		getResponseOK(w, r, weather)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, weather *entities.Weather) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Weather:  weather,
	})
}
