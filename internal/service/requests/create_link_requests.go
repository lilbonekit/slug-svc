package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/lilbonekit/slug-svc/internal/service/domain"
)

type CreateLinkRequest struct {
	TargetURL string `json:"target_url"`
	Slug      string `json:"slug,omitempty"`
	TTL       *int64 `json:"ttl,omitempty"`
}

type CreateLinkResponse struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

// Bind decodes + validates JSON body
func (r *CreateLinkRequest) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return errors.New("invalid JSON body")
	}

	r.TargetURL = strings.TrimSpace(r.TargetURL)
	if r.TargetURL == "" {
		return errors.New("target_url is required")
	}

	if err := domain.ValidateTargetURL(r.TargetURL); err != nil {
		return err
	}

	if r.Slug != "" {
		r.Slug = strings.TrimSpace(r.Slug)
	}

	return nil
}
