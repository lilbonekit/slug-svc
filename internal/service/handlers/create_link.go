package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"

	"github.com/lilbonekit/slug-svc/internal/service/domain"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"github.com/lilbonekit/slug-svc/internal/service/slugid"
)

type createReq struct {
	TargetURL string `json:"target_url"`
	Slug      string `json:"slug,omitempty"`
	TTL       *int64 `json:"ttl,omitempty"`
}

type createResp struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

func (h *Handlers) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req createReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ape.RenderErr(w, newError(http.StatusBadRequest, "invalid_json", err.Error()))
		return
	}
	if err := domain.ValidateTargetURL(req.TargetURL); err != nil {
		ape.RenderErr(w, newError(http.StatusUnprocessableEntity, "invalid_target_url", err.Error()))
		return
	}

	slug := strings.TrimSpace(req.Slug)
	if slug == "" {
		slug = slugid.Generate(6)
	}

	created, err := h.LinksRepo.Create(r.Context(), repo.Link{
		Slug:      slug,
		TargetURL: req.TargetURL,
		TTL:       req.TTL,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create link")
		ape.RenderErr(w, newError(http.StatusInternalServerError, "internal_error", "failed to create"))
		return
	}

	base := strings.TrimRight(h.BaseURL, "/")
	ape.Render(w, createResp{
		Slug: created.Slug,
		URL:  base + "/" + created.Slug,
	})
}
