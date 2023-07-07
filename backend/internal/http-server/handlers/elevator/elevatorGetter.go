package elevator

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type getResponse struct {
	resp.Response
	Elevator *entities.Elevator `json:"elevator"`
}

type getAllResponse struct {
	resp.Response
	Elevators []*entities.Elevator `json:"elevators"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=elevatorSaver
type Getter interface {
	GetAll() ([]*entities.Elevator, error)
	Get(id string) (*entities.Elevator, error)
	GetByBuilding(buildingId int) ([]*entities.Elevator, error)
}

func GetAll(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.elevator.GetAll"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		elevators, err := getter.GetAll()
		if err != nil {
			log.Error("failed to get elevators", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get elevators"))
			return
		}
		getAllResponseOK(w, r, elevators)
	}
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.elevator.Get"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id := chi.URLParam(r, "id")

		elevator, err := getter.Get(id)
		if err != nil {
			log.Error("failed to get elevator", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get elevator"))
			return
		}
		getResponseOK(w, r, elevator)
	}
}

func GetByBuilding(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.elevator.GetByBuilding"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		buildingId, err := strconv.Atoi(chi.URLParam(r, "building_id"))
		if err != nil {
			log.Error("failed to convert building_id", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to convert building_id"))
			return
		}
		elevator, err := getter.GetByBuilding(buildingId)
		if err != nil {
			log.Error("failed to get elevator", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get elevator"))
			return
		}
		getAllResponseOK(w, r, elevator)
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, elevators []*entities.Elevator) {
	render.JSON(w, r, getAllResponse{
		Response:  resp.OK(),
		Elevators: elevators,
	})
}
func getResponseOK(w http.ResponseWriter, r *http.Request, elevator *entities.Elevator) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Elevator: elevator,
	})
}
