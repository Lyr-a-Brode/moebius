package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Lyr-a-Brode/moebius/blob-store/api"
	"github.com/Lyr-a-Brode/moebius/blob-store/metrics"
	"github.com/Lyr-a-Brode/moebius/blob-store/service"
	"github.com/Lyr-a-Brode/moebius/blob-store/store"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.StandardLogger()
	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999",
	})

	cfg, err := parseConfig()
	if err != nil {
		log.WithError(err).Fatal("Unable to parse config")
	}

	if cfg.App.EnableDebug {
		logger.SetLevel(log.DebugLevel)
	}

	blobStore, err := createStore(cfg.Store)
	if err != nil {
		log.WithError(err).Fatal("Unable to create blob store")
	}

	apiServer := &http.Server{
		Addr:    cfg.App.Address,
		Handler: api.NewRouter(api.NewHandlers(service.NewStoreService(blobStore)), cfg.App.EnableDebug),
	}

	metricsServer := &http.Server{
		Addr:    cfg.App.MetricsAddress,
		Handler: metrics.NewRouter(),
	}

	go func() {
		log.WithField("address", cfg.App.MetricsAddress).Info("Starting metrics server")

		if err := runServer(metricsServer); err != nil {
			log.WithError(err).Fatal("Failed to start metrics server")
		}
	}()

	go handleShutdown(apiServer, metricsServer, cfg.App.ShutdownTimeout)

	log.WithField("address", cfg.App.Address).Info("Starting API server")

	if err := runServer(apiServer); err != nil {
		log.WithError(err).Fatal("Failed to start API server")
	}
}

func handleShutdown(apiServer *http.Server, metricsServer *http.Server, shutdownTimeout time.Duration) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-c

	log.WithField("signal", sig.String()).
		Info("Exit signal received. Starting shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	var g errgroup.Group

	g.Go(func() error { return apiServer.Shutdown(ctx) })
	g.Go(func() error { return metricsServer.Shutdown(ctx) })

	if err := g.Wait(); err != nil {
		log.WithError(err).Error("Unable to stop web server gracefully")
	}
}

func runServer(s *http.Server) error {
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func createStore(opts StoreOpts) (service.Store, error) {
	switch opts.Type {
	case "file":
		return store.NewFileStore(opts.Path), nil
	default:
		return nil, fmt.Errorf("unknown store type %s", opts.Type)
	}
}
