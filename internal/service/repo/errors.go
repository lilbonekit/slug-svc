package repo

import "errors"

var (
	ErrSlugExists = errors.New("slug already exists")
)
