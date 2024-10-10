package vehicle

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(router *chi.Mux, handler *handler) {
	router.Route("/vehicles", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Post("/", handler.Create)
	})
}
