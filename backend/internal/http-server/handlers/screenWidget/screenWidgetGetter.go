package screenWidget

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	ScreenWidgets []*entities.ScreenWidget `json:"screen_widgets"`
}

type Getter interface {
	Get(screenId string) ([]*entities.ScreenWidget, error)
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screenWidget.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		screenId := chi.URLParam(r, "screen_id")

		screen, err := getter.Get(screenId)
		if err != nil {
			log.Error("failed to get screenWidget", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get screenWidget"))

			return
		}

		getResponseOK(w, r, screen)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, screenWidgets []*entities.ScreenWidget) {
	render.JSON(w, r, getResponse{
		Response:      resp.OK(),
		ScreenWidgets: screenWidgets,
	})
}
