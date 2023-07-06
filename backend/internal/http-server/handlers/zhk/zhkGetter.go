package zhk

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/services/ujin/complexService"
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
	Zhk *entities.Zhk `json:"zhk"`
}

type getAllResponse struct {
	resp.Response
	Zhks *[]entities.Zhk `json:"zhks"`
}

type ComplexService interface {
	GetComplexes() (*complexService.Complex, error)
}

func GetAll(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.zhk.GetAll"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		complex, err := complexService.GetComplexes()
		if err != nil {
			log.Error("failed to get complexes", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get complexes"))
			return
		}

		var complexEntities []entities.Zhk
		complexMap := make(map[int]string)
		for _, building := range complex.Data.Buildings {
			complexMap[building.Complex.Id] = building.Complex.Title
		}

		for key, value := range complexMap {
			var complexEntity entities.Zhk
			complexEntity.Id = key
			complexEntity.Name = value
			complexEntities = append(complexEntities, complexEntity)
		}
		getAllResponseOK(w, r, &complexEntities)
	}
}

func Get(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.zhk.Get"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to convert id", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to convert id"))
			return
		}

		complex, err := complexService.GetComplexes()
		if err != nil {
			log.Error("failed to get complexes", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get complexes"))
			return
		}

		for _, building := range complex.Data.Buildings {
			if building.Id == id {
				var complexEntity entities.Zhk
				complexEntity.Id = building.Id
				complexEntity.Name = building.BuildingInfo.Title
				getResponseOK(w, r, &complexEntity)
			}
		}

		log.Error("failed to get complexes", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Error("failed to get complexes"))
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, zhks *[]entities.Zhk) {
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
