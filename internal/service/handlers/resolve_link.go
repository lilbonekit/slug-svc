package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (h *Handlers) ResolveLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		ape.RenderErr(w, newError(http.StatusBadRequest, "empty_slug", "slug is required"))
		return
	}

	link, err := h.LinksRepo.GetBySlug(r.Context(), slug)
	if err != nil {
		ape.RenderErr(w, newError(http.StatusNotFound, "not_found", "link not found"))
		return
	}

	if link.TTL != nil {
		expiredAt := link.CreatedAt.Add(time.Duration(*link.TTL) * time.Second)
		if time.Now().After(expiredAt) {
			ape.RenderErr(w, newError(http.StatusGone, "expired", "link expired"))
			return
		}
	}

	w.Header().Set("Location", link.TargetURL)
	w.WriteHeader(http.StatusFound)
}
