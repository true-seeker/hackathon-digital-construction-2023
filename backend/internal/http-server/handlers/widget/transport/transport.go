package transport

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
	Transport *[]entities.Transport `json:"transport"`
}

type Getter interface {
	GetTransport() (*[]entities.Transport, error)
}

func GetTransport(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transport, err := getter.GetTransport()

		if err != nil {
			log.Error("failed to get transport", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to get transport"))
			return
		}

		getResponseOK(w, r, transport)
	}
}

func getResponseOK(w http.ResponseWriter, r *http.Request, transport *[]entities.Transport) {
	render.JSON(w, r, getResponse{
		Response:  resp.OK(),
		Transport: transport,
	})
}
