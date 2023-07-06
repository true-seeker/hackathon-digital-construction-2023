package screen

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
	Name       string `json:"name,omitempty" validate:"required"`
	ElevatorId string `json:"elevator_id" validate:"required"`
	X          int    `json:"x" validate:"required"`
	Y          int    `json:"y" validate:"required"`
}

type saveResponse struct {
	resp.Response
	Screen *entities.Screen `json:"screen"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=screenSaver
type Saver interface {
	New(req *SaveRequest) (*entities.Screen, error)
}

func New(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.screen.New"

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

		screen, err := saver.New(&req)
		if err != nil {
			log.Error("failed to add screen", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add screen"))

			return
		}

		log.Info("screen added", slog.String("id", screen.Id))

		saverResponseOK(w, r, screen)
	}
}

func saverResponseOK(w http.ResponseWriter, r *http.Request, screen *entities.Screen) {
	render.JSON(w, r, saveResponse{
		Response: resp.OK(),
		Screen:   screen,
	})
}
