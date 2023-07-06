package building

import (
	"backend/internal/domain/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"

	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
)

type SaveRequest struct {
	Name      string  `json:"name,omitempty" validate:"required"`
	Address   string  `json:"address,omitempty" validate:"required"`
	ZhkId     string  `json:"zhk_id,omitempty" validate:"required"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type saveResponse struct {
	resp.Response
	Building *entities.Building `json:"building"`
}

type Location struct {
	Data []LocationData `json:"data"`
}

type LocationData struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=buildingSaver
type Saver interface {
	New(req *SaveRequest) (*entities.Building, error)
}

func New(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SaveRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode SaveRequest body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode SaveRequest"))

			return
		}

		log.Info("SaveRequest body decoded", slog.Any("SaveRequest", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid SaveRequest", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}
		location := getLocationByAddress(req.Address)
		if len(location.Data) > 0 {
			req.Latitude = location.Data[0].Latitude
			req.Longitude = location.Data[0].Longitude
		}
		building, err := saver.New(&req)
		if err != nil {
			log.Error("failed to add building", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add building"))

			return
		}

		log.Info("building added", slog.Int("id", building.Id))

		saverResponseOK(w, r, building)
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

func saverResponseOK(w http.ResponseWriter, r *http.Request, building *entities.Building) {
	render.JSON(w, r, saveResponse{
		Response: resp.OK(),
		Building: building,
	})
}
