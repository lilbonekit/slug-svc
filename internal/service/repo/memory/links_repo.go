package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/lilbonekit/slug-svc/internal/service/repo"
	"gitlab.com/distributed_lab/logan/v3"
)

type linksRepo struct {
	mu    sync.RWMutex
	store map[string]repo.Link
}

func (r *linksRepo) StartTTLWatcher(ctx context.Context, log *logan.Entry, interval time.Duration) {
	// TODO: Implement if needed
	panic("unimplemented")
}

func New() repo.LinksRepo {
	return &linksRepo{store: make(map[string]repo.Link)}
}

func (r *linksRepo) Create(ctx context.Context, l repo.Link) (repo.Link, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[l.Slug]; exists {
		return repo.Link{}, errors.New("slug already exists")
	}
	l.CreatedAt = time.Now()
	r.store[l.Slug] = l
	return l, nil
}

func (r *linksRepo) GetBySlug(ctx context.Context, slug string) (repo.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	l, ok := r.store[slug]
	if !ok {
		return repo.Link{}, errors.New("not found")
	}
	return l, nil
}
