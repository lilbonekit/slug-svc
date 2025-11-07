package service

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lilbonekit/slug-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

const (
	servicePath = "/integrations/slug-svc"
	v1Path      = "/v1"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(handlers.CtxLog(s.log)),
	)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.Route(servicePath, func(r chi.Router) {
		r.Route(v1Path, func(r chi.Router) {
			r.Post("/links", s.h.CreateLink)
		})
		r.Get("/{slug}", s.h.ResolveLink)
	})

	return r
}
