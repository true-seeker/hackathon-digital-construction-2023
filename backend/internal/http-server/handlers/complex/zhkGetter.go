package complex

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
	Complex *entities.Complex `json:"complex"`
}

type getAllResponse struct {
	resp.Response
	Complexes *[]entities.Complex `json:"complexes"`
}

type ComplexService interface {
	GetComplexes() (*complexService.Complex, error)
}

func GetAll(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.complex.GetAll"
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

		var complexEntities []entities.Complex
		complexMap := make(map[int]string)
		for _, building := range complex.Data.Buildings {
			complexMap[building.Complex.Id] = building.Complex.Title
		}

		for key, value := range complexMap {
			var complexEntity entities.Complex
			complexEntity.Id = key
			complexEntity.Name = value
			complexEntities = append(complexEntities, complexEntity)
		}
		getAllResponseOK(w, r, &complexEntities)
	}
}

func Get(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.complex.Get"
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
				var complexEntity entities.Complex
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

func getAllResponseOK(w http.ResponseWriter, r *http.Request, complexes *[]entities.Complex) {
	render.JSON(w, r, getAllResponse{
		Response:  resp.OK(),
		Complexes: complexes,
	})
}
func getResponseOK(w http.ResponseWriter, r *http.Request, complex *entities.Complex) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		Complex:  complex,
	})
}
