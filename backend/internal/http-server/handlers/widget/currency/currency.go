package currency

import (
	"backend/internal/domain/entities"
	resp "backend/internal/lib/api/response"
	"backend/internal/lib/logger/sl"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type getResponse struct {
	resp.Response
	Currencies *[]entities.Currency `json:"currencies"`
}

type Getter interface {
	GetCurrency() (*[]entities.Currency, error)
}

func GetCurrencies(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO weather by region
		currencies, err := getter.GetCurrency()

		if err != nil {
			log.Error("failed to get currencies", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get currencies"))

			return
		}

		getResponseOK(w, r, currencies)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, currencies *[]entities.Currency) {
	render.JSON(w, r, getResponse{
		Response:   resp.OK(),
		Currencies: currencies,
	})
}
