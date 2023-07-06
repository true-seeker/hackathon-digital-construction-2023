package building

import (
	"backend/internal/domain/entities"
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

func saverResponseOK(w http.ResponseWriter, r *http.Request, building *entities.Building) {
	render.JSON(w, r, saveResponse{
		Response: resp.OK(),
		Building: building,
	})
}
