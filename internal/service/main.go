package service

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/lilbonekit/slug-svc/internal/config"
	"github.com/lilbonekit/slug-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	deps     deps
	h        *handlers.Handlers
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	d := buildDeps(cfg)

	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		deps:     d,
		h:        handlers.New(d.Links(), d.BaseURL()),
	}
}

func Run(cfg config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	srv := newService(cfg)

	watcherCfg, err := cfg.Getter().GetStringMap("watcher")
	if err != nil {
		cfg.Log().WithError(err).Warn("failed to get watcher config, using default 1m")
	}

	intervalStr := "1m"
	if raw, ok := watcherCfg["interval"].(string); ok && raw != "" {
		intervalStr = raw
	}

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		cfg.Log().WithError(err).Warn("invalid watcher interval, using default 1m")
		interval = time.Minute
	}

	srv.deps.Links().StartTTLWatcher(ctx, cfg.Log(), interval)
	cfg.Log().WithField("interval", interval).Info("TTL watcher started")

	srv.log.Info("Service started")
	r := srv.router()
	httpSrv := &http.Server{
		Addr:    srv.listener.Addr().String(),
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		srv.log.Info("Shutting down HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpSrv.Shutdown(shutdownCtx); err != nil {
			srv.log.WithError(err).Error("HTTP server shutdown failed")
		} else {
			srv.log.Info("HTTP server stopped gracefully")
		}
	}()

	if err := httpSrv.Serve(srv.listener); err != nil && err != http.ErrServerClosed {
		panic(errors.Wrap(err, "failed to serve HTTP"))
	}
}
