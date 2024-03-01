package storage

import (
	"github.com/go-chi/chi/v5"
	"go-musthave-shortener-tpl/internal/app/compress"
	"go-musthave-shortener-tpl/internal/app/logger"
)

func LinkRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(compress.GzipMiddleware())
	r.Use(logger.Logger())
	r.Route("/", func(r chi.Router) {
		r.Post("/", CreateRedirect)
		r.Get("/{shortlink}", Redirect)
		r.Route("/api", func(r chi.Router) {
			r.Post("/shorten", APIShorten)
		})
	})
	return r
}
