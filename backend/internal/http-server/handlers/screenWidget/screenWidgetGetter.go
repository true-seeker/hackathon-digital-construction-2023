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

		screenWidgets, err := getter.Get(screenId)
		if err != nil {
			log.Error("failed to get screenWidget", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get screenWidget"))
			return
		}

		//
		if len(screenWidgets) == 0 {
			screenWidgets = hardCodeData()
		}

		getResponseOK(w, r, screenWidgets)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, screenWidgets []*entities.ScreenWidget) {
	render.JSON(w, r, getResponse{
		Response:      resp.OK(),
		ScreenWidgets: screenWidgets,
	})
}

func hardCodeData() []*entities.ScreenWidget {
	var screenWidgets []*entities.ScreenWidget
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "63baeddd-2a07-4f71-aa19-62ecbae26429",
		X:    1,
		Y:    5,
		W:    9,
		H:    1,
		MinW: 9,
		MinH: 1,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "d30bb91e-6718-4380-a196-9b791b26280d",
		X:    1,
		Y:    2,
		W:    7,
		H:    3,
		MinW: 7,
		MinH: 3,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "d6e2f387-a6ea-471b-96c3-d46a0e7c796d",
		X:    1,
		Y:    0,
		W:    7,
		H:    2,
		MinW: 3,
		MinH: 2,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "b71bef49-574e-4354-867c-ca77794172be",
		X:    1,
		Y:    12,
		W:    9,
		H:    3,
		MinW: 6,
		MinH: 3,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "e6b16a02-3d14-4185-b02b-ef1c3035f159",
		X:    1,
		Y:    6,
		W:    9,
		H:    6,
		MinW: 6,
		MinH: 6,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "070f62e1-dad3-454c-b89f-78df02df1039",
		X:    8,
		Y:    0,
		W:    4,
		H:    5,
		MinW: 4,
		MinH: 5,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "61493b97-7d24-4957-9d0a-3548f456374f",
		X:    0,
		Y:    0,
		W:    1,
		H:    4,
		MinW: 1,
		MinH: 1,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "e953c6b2-ce4d-42a1-b1b0-7a264172b1a2",
		X:    10,
		Y:    5,
		W:    2,
		H:    10,
		MinW: 1,
		MinH: 1,
	})
	screenWidgets = append(screenWidgets, &entities.ScreenWidget{
		I:    "7e551a5b-ff79-4c4e-81c7-2697478d6b54",
		X:    0,
		Y:    4,
		W:    1,
		H:    11,
		MinW: 1,
		MinH: 11,
	})
	return screenWidgets
}
