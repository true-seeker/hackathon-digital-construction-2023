package screen

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
	Screen *entities.Screen `json:"screen"`
}

type getAllResponse struct {
	resp.Response
	Screens []*entities.Screen `json:"screens"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=screenSaver
type Getter interface {
	GetAll() ([]*entities.Screen, error)
	Get(id string) (*entities.Screen, error)
	GetByElevator(elevatorId string) ([]*entities.Screen, error)
}

func GetAll(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screen.GetAll"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		screens, err := getter.GetAll()
		if err != nil {
			log.Error("failed to get screen", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get screen"))
			return
		}
		getAllResponseOK(w, r, screens)
	}
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screen.Get"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id := chi.URLParam(r, "id")

		screen, err := getter.Get(id)
		if err != nil {
			log.Error("failed to get screen", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get screen"))
			return
		}
		getResponseOK(w, r, screen)
	}
}

func GetByElevator(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.screen.GetByElevator"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		elevatorId := chi.URLParam(r, "elevator_id")

		screen, err := getter.GetByElevator(elevatorId)
		if err != nil {
			log.Error("failed to get screen", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get screen"))
			return
		}
		getAllResponseOK(w, r, screen)
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, screens []*entities.Screen) {
	render.JSON(w, r, getAllResponse{
		Response: resp.OK(),
		Screens:  screens,
	})
}
func getResponseOK(w http.ResponseWriter, r *http.Request, screen *entities.Screen) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Screen:   screen,
	})
}
