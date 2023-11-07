package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/ono5/myGoWebApplication/pkg/config"
	"github.com/ono5/myGoWebApplication/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
