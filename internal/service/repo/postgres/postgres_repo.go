package postgres

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"gitlab.com/distributed_lab/kit/pgdb"
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
