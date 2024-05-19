package main

import (
	"github.com/MatiasWebDevCoder4517/breakfast_reservations/pkg/config"
	"github.com/MatiasWebDevCoder4517/breakfast_reservations/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
