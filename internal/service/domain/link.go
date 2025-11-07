package domain

import (
	"errors"
	"net/url"
)

func ValidateTargetURL(u string) error {
	p, err := url.Parse(u)
	if err != nil || p.Scheme == "" || p.Host == "" {
		return errors.New("invalid URL")
	}
	return nil
}
