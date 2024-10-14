package vehicle

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(router *chi.Mux, handler *handler) {
	router.Route("/vehicles", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/color/{color}/year/{year}", handler.GetByColorAndYear)
		r.Post("/", handler.Create)
	})
}
