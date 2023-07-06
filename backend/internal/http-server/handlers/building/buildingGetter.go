package building

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
	Building *entities.Building `json:"building"`
}

type getAllResponse struct {
	resp.Response
	Buildings *[]entities.Building `json:"buildings"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=buildingSaver
type Getter interface {
	GetAll() ([]*entities.Building, error)
	Get(id string) (*entities.Building, error)
}

type ComplexService interface {
	GetComplexes() (complexService.Complex, error)
}

func GetAll(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		complex, err := complexService.GetComplexes()
		if err != nil {
			log.Error("failed to get complexes", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get complexes"))

			return
		}

		var buildingEntities []entities.Building
		buildingMap := make(map[int]entities.Building)
		for _, building := range complex.Data.Buildings {
			buildingMap[building.Id] = entities.Building{
				Id:      building.Id,
				Name:    building.BuildingInfo.Title,
				Address: building.BuildingInfo.Address.FullAddress,
				ZhkId:   building.Complex.Id,
			}
		}

		for _, value := range buildingMap {
			buildingEntities = append(buildingEntities, value)
		}
		getAllResponseOK(w, r, &buildingEntities)
	}
}

func Get(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.building.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to convert id", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to convert id"))

			return
		}
		//
		//building, err := getter.Get(id)
		//if err != nil {
		//	log.Error("failed to get building", sl.Err(err))
		//
		//	render.JSON(w, r, resp.Error("failed to get building"))
		//
		//	return
		//}
		//getResponseOK(w, r, building)

		complex, err := complexService.GetComplexes()
		if err != nil {
			log.Error("failed to get complexes", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get complexes"))

			return
		}

		for _, building := range complex.Data.Buildings {
			if building.Id == id {
				getResponseOK(w, r, &entities.Building{
					Id:      building.Id,
					Name:    building.BuildingInfo.Title,
					Address: building.BuildingInfo.Address.FullAddress,
					ZhkId:   building.Complex.Id,
				})
				break
			}
		}
	}
}

func GetByZhk(log *slog.Logger, complexService ComplexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.buildings.Get"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		zhk_id, err := strconv.Atoi(chi.URLParam(r, "zhk_id"))
		if err != nil {
			log.Error("failed to convert zhk", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to convert zhk"))

			return
		}

		complex, err := complexService.GetComplexes()
		if err != nil {
			log.Error("failed to get complexes", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get complexes"))

			return
		}

		var buildingEntities []entities.Building
		buildingMap := make(map[int]entities.Building)
		for _, building := range complex.Data.Buildings {
			if building.Complex.Id == zhk_id {
				buildingMap[building.Id] = entities.Building{
					Id:      building.Id,
					Name:    building.BuildingInfo.Title,
					Address: building.BuildingInfo.Address.FullAddress,
					ZhkId:   zhk_id,
				}
			}
		}

		for _, value := range buildingMap {
			buildingEntities = append(buildingEntities, value)
		}
		getAllResponseOK(w, r, &buildingEntities)
	}
}

func getAllResponseOK(w http.ResponseWriter, r *http.Request, buildings *[]entities.Building) {
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
