package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/lilbonekit/slug-svc/internal/service/domain"
	"github.com/lilbonekit/slug-svc/resources"
)

func Bind(r *http.Request) (*resources.CreateLinkRequest, error) {
	var req resources.CreateLinkRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid JSON body: %w", err)
	}

	attrs := req.Data.Attributes

	if attrs.TargetUrl == nil || strings.TrimSpace(*attrs.TargetUrl) == "" {
		return nil, errors.New("target_url is required")
	}

	url := strings.TrimSpace(*attrs.TargetUrl)
	if err := domain.ValidateTargetURL(url); err != nil {
		return nil, err
	}
	attrs.TargetUrl = &url

	if attrs.Slug != nil {
		slug := strings.TrimSpace(*attrs.Slug)
		attrs.Slug = &slug
	}

	req.Data.Attributes = attrs
	return &req, nil
}
