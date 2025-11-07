package service

import (
	"github.com/lilbonekit/slug-svc/internal/config"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	mem "github.com/lilbonekit/slug-svc/internal/service/repo/memory"
	pgrepo "github.com/lilbonekit/slug-svc/internal/service/repo/postgres"
	"gitlab.com/distributed_lab/kit/kv"
)

type deps struct {
	links   repo.LinksRepo
	baseURL string
}

func (d deps) Links() repo.LinksRepo { return d.links }
func (d deps) BaseURL() string       { return d.baseURL }

func buildDeps(cfg config.Config) deps {
	getter := cfg.Getter()
	section := kv.MustGetStringMap(getter, "shortener")

	base, _ := section["base_url"].(string)
	if base == "" {
		base = "http://localhost:8000"
	}

	storage := "memory"
	if v, ok := section["storage"].(string); ok && v != "" {
		storage = v
	}

	var links repo.LinksRepo
	switch storage {
	case "pg", "postgres", "postgresql":
		links = pgrepo.New(cfg.DB())
	default:
		links = mem.New()
	}

	return deps{
		links:   links,
		baseURL: base,
	}
}
