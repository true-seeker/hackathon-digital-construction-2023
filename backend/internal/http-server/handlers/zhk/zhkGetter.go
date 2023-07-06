package zhk

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
	Zhk *entities.Zhk `json:"zhk"`
}

type getAllResponse struct {
	resp.Response
	Zhks []*entities.Zhk `json:"zhks"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=zhkSaver
type Getter interface {
	GetAll() ([]*entities.Zhk, error)
	Get(id string) (*entities.Zhk, error)
}

func GetAll(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.zhk.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		zhks, err := getter.GetAll()
		if err != nil {
			log.Error("failed to get zhks", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get zhks"))

			return
		}
		getAllResponseOK(w, r, zhks)
	}
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.zhk.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id := chi.URLParam(r, "id")

		zhk, err := getter.Get(id)
		if err != nil {
			log.Error("failed to get zhk", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get zhk"))

			return
		}
		getResponseOK(w, r, zhk)
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, zhks []*entities.Zhk) {
	render.JSON(w, r, getAllResponse{
		Response: resp.OK(),
		Zhks:     zhks,
	})
}
func getResponseOK(w http.ResponseWriter, r *http.Request, zhk *entities.Zhk) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Zhk:      zhk,
	})
}
