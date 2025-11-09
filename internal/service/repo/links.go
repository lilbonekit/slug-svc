package repo

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
)

type Link struct {
	Slug      string    `db:"slug"       json:"slug"`
	TargetURL string    `db:"target_url" json:"target_url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	TTL       *int64    `db:"ttl"        json:"ttl"`
}

type LinksRepo interface {
	Create(ctx context.Context, l Link) (Link, error)
	GetBySlug(ctx context.Context, slug string) (Link, error)
	StartTTLWatcher(ctx context.Context, log *logan.Entry, interval time.Duration)
}
