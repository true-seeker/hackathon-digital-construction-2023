package elevator

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
	BuildingId int    `json:"building_id" validate:"required"`
}

type saveResponse struct {
	resp.Response
	Elevator *entities.Elevator `json:"elevator"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=elevatorSaver
type Saver interface {
	New(req *SaveRequest) (*entities.Elevator, error)
}

func New(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.elevator.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SaveRequest

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
			log.Error("invalid SaveRequest", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		elevator, err := saver.New(&req)
		if err != nil {
			log.Error("failed to add elevator", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to add elevator"))
			return
		}

		log.Info("elevator added", slog.String("id", elevator.Id))

		saverResponseOK(w, r, elevator)
	}
}

func saverResponseOK(w http.ResponseWriter, r *http.Request, elevator *entities.Elevator) {
	render.JSON(w, r, saveResponse{
		Response: resp.OK(),
		Elevator: elevator,
	})
}
