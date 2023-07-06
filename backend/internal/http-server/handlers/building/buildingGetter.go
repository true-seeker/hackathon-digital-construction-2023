package building

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
	Building *entities.Building `json:"building"`
}

type getAllResponse struct {
	resp.Response
	Buildings []*entities.Building `json:"buildings"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=buildingSaver
type Getter interface {
	GetAll() ([]*entities.Building, error)
	Get(id string) (*entities.Building, error)
	GetByZhk(id string) (*entities.Building, error)
}

func GetAll(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		buildings, err := getter.GetAll()
		if err != nil {
			log.Error("failed to get buildings", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get buildings"))

			return
		}
		getAllResponseOK(w, r, buildings)
	}
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id := chi.URLParam(r, "id")

		building, err := getter.Get(id)
		if err != nil {
			log.Error("failed to get building", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get building"))

			return
		}
		getResponseOK(w, r, building)
	}
}

func GetByZhk(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.zhk.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		zhk_id := chi.URLParam(r, "zhk_id")

		zhk, err := getter.GetByZhk(zhk_id)
		if err != nil {
			log.Error("failed to get zhk", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get zhk"))

			return
		}
		getResponseOK(w, r, zhk)
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, buildings []*entities.Building) {
	render.JSON(w, r, getAllResponse{
		Response:  resp.OK(),
		Buildings: buildings,
	})
}
func getResponseOK(w http.ResponseWriter, r *http.Request, building *entities.Building) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Building: building,
	})
}
