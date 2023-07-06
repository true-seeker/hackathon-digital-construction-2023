package main

import (
	"backend/internal/http-server/handlers/building"
	"backend/internal/http-server/handlers/elevator"
	screen "backend/internal/http-server/handlers/screen"
	"backend/internal/http-server/handlers/zhk"
	"backend/internal/storage/postgres"
	"backend/internal/storage/postgres/repository"
	"fmt"
	"github.com/go-chi/cors"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"

	"backend/internal/config"
	mwLogger "backend/internal/http-server/middleware/logger"
	"backend/internal/lib/logger/handlers/slogpretty"
	"backend/internal/lib/logger/sl"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Env)
	log := setupLogger(cfg.Env)

	log.Info(
		"starting backend",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)

	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	buildingRepository := repository.NewBuildingRepository(storage.GetDb())
	elevatorRepository := repository.NewElevatorRepository(storage.GetDb())
	screenRepository := repository.NewScreenRepository(storage.GetDb())
	zhkRepository := repository.NewZhkRepository(storage.GetDb())

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/api", func(r chi.Router) {
		//r.Use(middleware.BasicAuth("hackathon-digital-construction-2023", map[string]string{
		//	cfg.HTTPServer.User: cfg.HTTPServer.Password,
		//}))
		r.Route("/buildings", func(r chi.Router) {
			r.Post("/", building.New(log, buildingRepository))
			r.Get("/", building.GetAll(log, buildingRepository))
			r.Get("/{id}", building.Get(log, buildingRepository))
			r.Get("/{building_id}/elevators", elevator.GetByBuilding(log, elevatorRepository))
			r.Put("/", building.Update(log, buildingRepository))
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

		r.Route("/zhks", func(r chi.Router) {
			r.Post("/", zhk.New(log, zhkRepository))
			r.Get("/", zhk.GetAll(log, zhkRepository))
			r.Get("/{id}", zhk.Get(log, zhkRepository))
			r.Get("/{zhk_id}/buildings", building.GetByZhk(log, buildingRepository))
			r.Put("/", zhk.Update(log, zhkRepository))
		})
	})

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
