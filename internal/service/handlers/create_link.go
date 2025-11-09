package handlers

import (
	"errors"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"

	"github.com/lilbonekit/slug-svc/internal/service/domain"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"github.com/lilbonekit/slug-svc/internal/service/requests"
	"github.com/lilbonekit/slug-svc/internal/service/slugid"
	"github.com/lilbonekit/slug-svc/resources"
)

func (h *Handlers) CreateLink(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Bind(r)
	if err != nil {
		ape.RenderErr(w, newError(http.StatusBadRequest, "invalid_request", err.Error()))
		return
	}

	attrs := req.Data.Attributes

	targetURL := strings.TrimSpace(*attrs.TargetUrl)
	if err := domain.ValidateTargetURL(targetURL); err != nil {
		ape.RenderErr(w, newError(http.StatusBadRequest, "invalid_url", err.Error()))
		return
	}

	slug := ""
	if attrs.Slug != nil && strings.TrimSpace(*attrs.Slug) != "" {
		slug = strings.TrimSpace(*attrs.Slug)
	} else {
		slug = slugid.Generate(6)
	}

	created, err := h.LinksRepo.Create(r.Context(), repo.Link{
		Slug:      slug,
		TargetURL: targetURL,
		TTL:       attrs.Ttl.Get(),
	})
	if err != nil {
		if errors.Is(err, repo.ErrSlugExists) {
			ape.RenderErr(w, newError(http.StatusConflict, "slug_exists", "slug already exists"))
			return
		}

		Log(r).WithError(err).Error("failed to create link")
		ape.RenderErr(w, newError(http.StatusInternalServerError, "internal_error", "failed to create link"))
		return
	}

	resp := resources.NewLink(resources.LinkAttributes{
		Slug:      created.Slug,
		TargetUrl: created.TargetURL,
		Ttl:       *resources.NewNullableInt64(created.TTL),
	})
	resp.Id = resources.PtrString(created.Slug)
	resp.Type = resources.PtrString("link")

	base := strings.TrimRight(h.BaseURL, "/")
	location := base + "/" + created.Slug

	w.Header().Set("Location", location)
	ape.Render(w, resp)
}
