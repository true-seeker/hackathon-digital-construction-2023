package market

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/services/ujin/market"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	MarketOffers *[]entities.MarketOffer `json:"news"`
}

type MarketService interface {
	GetMarketOffers() (*market.MarketOffers, error)
}

func GetMarketOffers(log *slog.Logger, marketService MarketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mf, err := marketService.GetMarketOffers()
		if err != nil {
			log.Error("failed to get market offers", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get market offers"))
			return
		}
		var marketOffers []entities.MarketOffer
		for _, offer := range mf.Data.Offers {
			var marketOfferEntity entities.MarketOffer
			marketOfferEntity.Title = offer.Title
			marketOfferEntity.ImageUrl = offer.Image.Url
			marketOfferEntity.Summary = offer.Summary
			marketOfferEntity.Id = offer.Id
			marketOffers = append(marketOffers, marketOfferEntity)
		}

		getResponseOK(w, r, &marketOffers)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, marketOffers *[]entities.MarketOffer) {
	render.JSON(w, r, getResponse{
		Response:     resp.OK(),
		MarketOffers: marketOffers,
	})
}
