package handlers

import (
	"errors"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/lilbonekit/slug-svc/internal/service/domain"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"github.com/lilbonekit/slug-svc/internal/service/requests"
	"github.com/lilbonekit/slug-svc/internal/service/slugid"
	"github.com/lilbonekit/slug-svc/resources"
)

func (h *Handlers) CreateLink(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Bind(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	attrs := req.Data.Attributes
	targetURL := strings.TrimSpace(*attrs.TargetUrl)

	if err := domain.ValidateTargetURL(targetURL); err != nil {
		ape.RenderErr(w, problems.BadRequest(errors.New("invalid target_url"))...)
		return
	}

	var slug string
	if attrs.Slug != nil && strings.TrimSpace(*attrs.Slug) != "" {
		slug = strings.TrimSpace(*attrs.Slug)
	} else {
		slug, err = slugid.Generate(6)
		if err != nil {
			Log(r).WithError(err).Error("failed to generate slug")
			ape.RenderErr(w, []*jsonapi.ErrorObject{problems.InternalError()}...)
			return
		}
	}

	created, err := h.LinksRepo.Create(r.Context(), repo.Link{
		Slug:      slug,
		TargetURL: targetURL,
		TTL:       attrs.Ttl.Get(),
	})

	if errors.Is(err, repo.ErrSlugExists) {
		validationErr := validation.Errors{
			"slug": errors.New("slug already exists"),
		}.Filter()

		if validationErr != nil {
			ape.RenderErr(w, problems.BadRequest(validationErr)...)
		} else {
			ape.RenderErr(w, problems.BadRequest(errors.New("slug already exists"))...)
		}
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
