package weather

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	Weather *entities.Weather `json:"weather"`
}

type Getter interface {
	GetWeather(building *entities.Building) (*entities.Weather, error)
}

type BuildingGetter interface {
	GetBuildingByScreenId(screenId string) (*entities.Building, error)
}

func GetWeather(log *slog.Logger, getter Getter, buildingGetter BuildingGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO weather by region
		screenId := chi.URLParam(r, "screen_id")
		building, err := buildingGetter.GetBuildingByScreenId(screenId)
		if err != nil {
			log.Error("failed to get building", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get building"))

			return
		}
		weather, err := getter.GetWeather(building)

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
