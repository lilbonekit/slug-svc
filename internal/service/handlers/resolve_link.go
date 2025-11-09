package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func (h *Handlers) ResolveLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		ape.RenderErr(w, problems.BadRequest(errors.New("slug is required"))...)
		return
	}

	link, err := h.LinksRepo.GetBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ape.RenderErr(w, []*jsonapi.ErrorObject{problems.NotFound()}...)
			return
		}

		Log(r).WithError(err).Error("failed to get link")
		ape.RenderErr(w, []*jsonapi.ErrorObject{problems.InternalError()}...)
		return
	}

	if link.TTL != nil {
		expiredAt := link.CreatedAt.Add(time.Duration(*link.TTL) * time.Second)
		if time.Now().After(expiredAt) {
			ape.RenderErr(w, []*jsonapi.ErrorObject{
				{
					Status: "410",
					Title:  "Gone",
					Detail: "link expired",
				},
			}...)
			return
		}
	}

	w.Header().Set("Location", link.TargetURL)
	w.WriteHeader(http.StatusFound)
}
