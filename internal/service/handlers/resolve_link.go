package handlers

import (
	"net/http"

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

	w.Header().Set("Location", link.TargetURL)
	w.WriteHeader(http.StatusFound)
}
