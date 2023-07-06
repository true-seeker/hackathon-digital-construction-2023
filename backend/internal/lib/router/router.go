package router

import (
	"backend/internal/config"
	"backend/internal/http-server/handlers/building"
	"backend/internal/http-server/handlers/elevator"
	"backend/internal/http-server/handlers/screen"
	"backend/internal/http-server/handlers/screenWidget"
	currencyWidget "backend/internal/http-server/handlers/widget/currency"
	marketWidget "backend/internal/http-server/handlers/widget/market"
	newsWidget "backend/internal/http-server/handlers/widget/news"
	weatherWidget "backend/internal/http-server/handlers/widget/weather"
	"backend/internal/http-server/handlers/zhk"
	"backend/internal/http-server/services/ujin/complexService"
	"backend/internal/http-server/services/ujin/market"
	"backend/internal/http-server/services/widget/currency"
	"backend/internal/http-server/services/widget/news"
	"backend/internal/http-server/services/widget/weather"
	"backend/internal/lib/logger/sl"
	"backend/internal/storage/postgres"
	"backend/internal/storage/postgres/repository"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
	"os"
)

func InitRoutes(router *chi.Mux, log *slog.Logger, cfg *config.Config) {
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	buildingRepository := repository.NewBuildingRepository(storage.GetDb())
	elevatorRepository := repository.NewElevatorRepository(storage.GetDb())
	screenRepository := repository.NewScreenRepository(storage.GetDb())
	screenWidgetRepository := repository.NewScreenWidgetRepository(storage.GetDb())

	weatherService := weather.NewWeatherYandexService()
	currencyService := currency.NewCurrencyService()
	complexSrvc := complexService.NewComplexService("ust-738975-7546d637ccea657af49f1eeecd9e4806")
	newsService := news.NewNewsService("ust-739005-c8bb190f92541a2c1e839b98cbc469c9")
	marketService := market.NewMarketService("ust-739005-c8bb190f92541a2c1e839b98cbc469c9")

	router.Route("/api", func(r chi.Router) {
		r.Route("/buildings", func(r chi.Router) {
			r.Get("/", building.GetAll(log, complexSrvc))
			r.Get("/{id}", building.Get(log, complexSrvc))
			r.Get("/{building_id}/elevators", elevator.GetByBuilding(log, elevatorRepository))
		})
		r.Route("/elevators", func(r chi.Router) {
			r.Post("/", elevator.New(log, elevatorRepository))
			r.Get("/", elevator.GetAll(log, elevatorRepository))
			r.Get("/{id}", elevator.Get(log, elevatorRepository))
			r.Get("/{elevator_id}/screens", screen.GetByElevator(log, screenRepository))
			r.Put("/", elevator.Update(log, elevatorRepository))
		})

		r.Route("/screens", func(r chi.Router) {
			r.Post("/", screen.New(log, screenRepository))
			r.Get("/", screen.GetAll(log, screenRepository))
			r.Get("/{id}", screen.Get(log, screenRepository))
			r.Put("/", screen.Update(log, screenRepository))
		})

		r.Route("/screen_widgets", func(r chi.Router) {
			r.Post("/", screenWidget.Save(log, screenWidgetRepository))
			r.Get("/{screen_id}", screenWidget.Get(log, screenWidgetRepository))
		})

		r.Route("/zhks", func(r chi.Router) {
			r.Get("/", zhk.GetAll(log, complexSrvc))
			r.Get("/{id}", zhk.Get(log, complexSrvc))
			r.Get("/{zhk_id}/buildings", building.GetByZhk(log, complexSrvc))
		})

		r.Route("/widgets", func(r chi.Router) {
			r.Get("/weather/{screen_id}", weatherWidget.GetWeather(log, weatherService, buildingRepository, complexSrvc))
			r.Get("/currency", currencyWidget.GetCurrencies(log, currencyService))
			r.Get("/news/{screen_id}", newsWidget.GetNews(log, newsService, buildingRepository))
			r.Get("/market", marketWidget.GetMarketOffers(log, marketService))
		})
	})

}
