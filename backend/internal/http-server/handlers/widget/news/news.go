package news

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/services/widget/news"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	News *[]entities.News `json:"news"`
}

type NewsService interface {
	GetNews(buildingId int) (*news.News, error)
}

type BuildingGetter interface {
	GetBuildingIdByScreenId(screenId string) (int, error)
}

func GetNews(log *slog.Logger, newsService NewsService, buildingGetter BuildingGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		screenId := chi.URLParam(r, "screen_id")

		buildingId, err := buildingGetter.GetBuildingIdByScreenId(screenId)
		if err != nil {
			log.Error("failed to get buildingId", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get buildingId"))

			return
		}

		ns, err := newsService.GetNews(buildingId)
		if err != nil {
			log.Error("failed to get news", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get news"))

			return
		}
		var newsEntities []entities.News
		for _, item := range ns.Data.Items {
			var newsEntity entities.News
			newsEntity.Title = item.Title
			newsEntity.Text = item.Text
			newsEntity.Date = item.Date
			newsEntities = append(newsEntities, newsEntity)
		}

		getResponseOK(w, r, &newsEntities)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, news *[]entities.News) {
	render.JSON(w, r, getResponse{
		Response: resp.OK(),
		News:     news,
	})
}
