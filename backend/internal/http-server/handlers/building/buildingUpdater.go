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

type UpdateRequest struct {
	Id      string `json:"id,omitempty" validate:"required"`
	Name    string `json:"name,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
}

type updateResponse struct {
	resp.Response
	Building *entities.Building `json:"building"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=buildingSaver
type Updater interface {
	Update(req *UpdateRequest) (*entities.Building, error)
}

func Update(log *slog.Logger, updater Updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.Update"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req UpdateRequest

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

		building, err := updater.Update(&req)
		if err != nil {
			log.Error("failed to add building", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add building"))

			return
		}

		log.Info("building added", slog.String("id", building.Id))

		saverResponseOK(w, r, building)
	}
}

func updaterResponseOK(w http.ResponseWriter, r *http.Request, building *entities.Building) {
	render.JSON(w, r, updateResponse{
		Response: resp.OK(),
		Building: building,
	})
}
