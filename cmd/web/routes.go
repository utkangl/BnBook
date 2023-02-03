package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/utkangl/GoWEB/packages/config"
	"github.com/utkangl/GoWEB/packages/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Recovers from possible panics and gives detaield information about what went wrong
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/book", handlers.Repo.GetBook)
	mux.Post("/book", handlers.Repo.PostBook)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/kings_suit", handlers.Repo.Kings_suit)
	mux.Get("/regular_room", handlers.Repo.Regular_room)
	mux.Get("/book-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/make_reservation", handlers.Repo.Make_reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
