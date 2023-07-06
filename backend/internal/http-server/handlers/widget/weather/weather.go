package weather

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/services/ujin/complexService"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
)

type getResponse struct {
	resp.Response
	Weather *entities.Weather `json:"weather"`
}

type Location struct {
	Data []LocationData `json:"data"`
}

type LocationData struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Getter interface {
	GetWeather(locationData LocationData) (*entities.Weather, error)
}

type BuildingGetter interface {
	GetBuildingIdByScreenId(screenId string) (int, error)
}

type ComplexService interface {
	GetBuildingById(buildingId int) (*complexService.BuildingInfo, error)
}

var weatherCache map[string]*entities.Weather

func GetWeather(log *slog.Logger, getter Getter, buildingGetter BuildingGetter, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if weatherCache == nil {
			weatherCache = make(map[string]*entities.Weather)
		}
		screenId := chi.URLParam(r, "screen_id")
		weather, ok := weatherCache[screenId]
		if ok {
			getResponseOK(w, r, weather)
			return
		}

		buildingId, err := buildingGetter.GetBuildingIdByScreenId(screenId)
		if err != nil {
			log.Error("failed to get buildingId", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get buildingId"))

			return
		}
		buildingInfo, err := complexService.GetBuildingById(buildingId)
		if err != nil {
			log.Error("failed to get buildingInfo", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get buildingInfo"))

			return
		}
		location := getLocationByAddress(buildingInfo.Address.FullAddress)
		var locationData LocationData
		if len(location.Data) > 0 {
			locationData = location.Data[0]
		}
		weather, err = getter.GetWeather(locationData)

		if err != nil {
			log.Error("failed to get weather", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get weather"))

			return
		}

		weatherCache[screenId] = weather
		getResponseOK(w, r, weather)
	}
}

func getLocationByAddress(address string) Location {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://api.positionstack.com/v1/forward?"+
		"access_key=ea9717bef2ef10ba7075c1f42e99be2a"+
		"& query="+address, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var location Location

	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return location
}

func getResponseOK(w http.ResponseWriter, r *http.Request, weather *entities.Weather) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Weather:  weather,
	})
}
