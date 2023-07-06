package screenWidget

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"net/http"
)

type SaveRequest struct {
	ScreenId      string                   `json:"screen_id" validate:"required"`
	ScreenWidgets *[]entities.ScreenWidget `json:"screen_widgets" validate:"required"`
}

type saveResponse struct {
	resp.Response
	ScreenId      string                   `json:"screen_id"`
	ScreenWidgets *[]entities.ScreenWidget `json:"screen_widgets"`
}

type Saver interface {
	Save(req *SaveRequest) (*[]entities.ScreenWidget, error)
}

func Save(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screenWidget.Save"

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

		screen, err := saver.Save(&req)
		if err != nil {
			log.Error("failed to add screenWidget", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add screenWidget"))

			return
		}

		saverResponseOK(w, r, screen)
	}
}

func saverResponseOK(w http.ResponseWriter, r *http.Request, screenWidgets *[]entities.ScreenWidget) {
	render.JSON(w, r, saveResponse{
		Response:      resp.OK(),
		ScreenWidgets: screenWidgets,
	})
}
