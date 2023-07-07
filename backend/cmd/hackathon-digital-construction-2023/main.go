package main

import (
	routerLib "backend/internal/lib/router"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/exp/slog"
	"net/http"

	"backend/internal/config"
	mwLogger "backend/internal/http-server/middleware/logger"
	"backend/internal/lib/logger/sl"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Env)
	log := sl.SetupLogger(cfg.Env)

	log.Info(
		"starting web",
		slog.String("env", cfg.Env),
	)

	router := chi.NewRouter()
	router.Use()
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
	router.Use(middleware.NoCache)

	routerLib.InitRoutes(router, log, cfg)

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.HTTPServer.Address, cfg.HTTPServer.Port),
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", err)
	}

	log.Error("server stopped")
}
