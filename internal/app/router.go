package server

import "github.com/go-chi/chi/v5"

func LinkRouter() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/", CreateRedirect)
		r.Get("/{shortlink}", Redirect)
	})
	return r
}
