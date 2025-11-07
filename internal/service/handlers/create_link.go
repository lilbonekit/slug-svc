package handlers

import (
	"errors"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"

	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"github.com/lilbonekit/slug-svc/internal/service/requests"
	"github.com/lilbonekit/slug-svc/internal/service/slugid"
)

func (h *Handlers) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateLinkRequest

	if err := req.Bind(r); err != nil {
		ape.RenderErr(w, newError(http.StatusBadRequest, "invalid_json", err.Error()))
		return
	}

	slug := req.Slug
	if slug == "" {
		slug = slugid.Generate(6)
	}

	created, err := h.LinksRepo.Create(r.Context(), repo.Link{
		Slug:      slug,
		TargetURL: req.TargetURL,
		TTL:       req.TTL,
	})
	if err != nil {
		if errors.Is(err, repo.ErrSlugExists) {
			ape.RenderErr(w, newError(http.StatusConflict, "slug_exists", "slug already exists"))
			return
		}

		Log(r).WithError(err).Error("failed to create link")
		ape.RenderErr(w, newError(http.StatusInternalServerError, "internal_error", "failed to create"))
		return
	}

	base := strings.TrimRight(h.BaseURL, "/")
	ape.Render(w, requests.CreateLinkResponse{
		Slug: created.Slug,
		URL:  base + "/" + created.Slug,
	})
}
