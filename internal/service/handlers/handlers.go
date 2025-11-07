package handlers

import "github.com/lilbonekit/slug-svc/internal/service/repo"

type Handlers struct {
	LinksRepo repo.LinksRepo
	BaseURL   string
}

func New(links repo.LinksRepo, baseURL string) *Handlers {
	return &Handlers{
		LinksRepo: links,
		BaseURL:   baseURL,
	}
}
