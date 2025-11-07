package handlers

import (
	"strconv"

	"github.com/google/jsonapi"
)

func newError(status int, title, detail string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  title,
		Detail: detail,
		Status: strconv.Itoa(status),
	}
}
