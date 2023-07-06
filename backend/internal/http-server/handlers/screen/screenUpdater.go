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

type UpdateRequest struct {
	Id         string `json:"id,omitempty" validate:"required"`
	Name       string `json:"name,omitempty" validate:"required"`
	ElevatorId string `json:"elevator_id" validate:"required"`
	X          int    `json:"x" validate:"required"`
	Y          int    `json:"y" validate:"required"`
}

type updateResponse struct {
	resp.Response
	Screen *entities.Screen `json:"screen"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=screenSaver
type Updater interface {
	Update(req *UpdateRequest) (*entities.Screen, error)
}

func Update(log *slog.Logger, updater Updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screen.Update"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req UpdateRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode SaveRequest body", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to decode SaveRequest"))
			return
		}

		log.Info("SaveRequest body decoded", slog.Any("SaveRequest", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			render.Status(r, http.StatusBadRequest)
			log.Error("invalid SaveRequest", sl.Err(err))
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		screen, err := updater.Update(&req)
		if err != nil {
			log.Error("failed to add screen", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to add screen"))
			return
		}

		log.Info("screen added", slog.String("id", screen.Id))

		updaterResponseOK(w, r, screen)
	}
}

func updaterResponseOK(w http.ResponseWriter, r *http.Request, screen *entities.Screen) {
	render.JSON(w, r, updateResponse{
		Response: resp.OK(),
		Screen:   screen,
	})
}
