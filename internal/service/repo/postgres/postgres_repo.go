package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
)

type linksRepo struct {
	db *pgdb.DB
	sb squirrel.StatementBuilderType
}

func New(db *pgdb.DB) repo.LinksRepo {
	return &linksRepo{
		db: db,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *linksRepo) Create(ctx context.Context, l repo.Link) (repo.Link, error) {
	q := r.sb.
		Insert("links").
		Columns("slug", "target_url", "ttl").
		Values(l.Slug, l.TargetURL, l.TTL).
		Suffix("RETURNING slug, target_url, created_at, ttl")

	var out repo.Link
	if err := r.db.GetContext(ctx, &out, q); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return repo.Link{}, repo.ErrSlugExists

		}
		return repo.Link{}, err
	}
	return out, nil
}

func (r *linksRepo) GetBySlug(ctx context.Context, slug string) (repo.Link, error) {
	q := r.sb.
		Select("slug", "target_url", "created_at", "ttl").
		From("links").
		Where(squirrel.Eq{"slug": slug}).
		Limit(1)

	var out repo.Link
	if err := r.db.GetContext(ctx, &out, q); err != nil {
		if err == sql.ErrNoRows {
			return repo.Link{}, err
		}
		return repo.Link{}, err
	}
	return out, nil
}

func (r *linksRepo) DeleteExpired(ctx context.Context) error {
	q := r.sb.
		Delete("links").
		Where("ttl IS NOT NULL").
		Where("(created_at + make_interval(secs => ttl)) < now()")

	return r.db.ExecContext(ctx, q)
}

func (r *linksRepo) StartTTLWatcher(ctx context.Context, log *logan.Entry, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := r.DeleteExpired(ctx); err != nil {
					log.WithError(err).Error("failed to delete expired links")
				} else {
					log.Info("expired links removed successfully")
				}
			case <-ctx.Done():
				ticker.Stop()
				log.Info("ttl watcher stopped")
				return
			}
		}
	}()
}
