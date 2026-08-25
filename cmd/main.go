package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/memap-project/memap-core/clean"
	"github.com/memap-project/memap-core/ns"
	"github.com/memap-project/memap/config"
	"github.com/memap-project/memap/logger"
	"github.com/memap-project/memap/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}
	if err = cfg.Validate(); err != nil {
		slog.Error("failed to validate config", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("config loaded")

	file, err := logger.Setup(cfg.Logger.LogPath)
	if err != nil {
		slog.Error("failed to setup logger", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer file.Close()

	manager := ns.NewNamespaceManager(&cfg.Core.Namespace)
	cleaner := clean.NewCleaner(ctx, cfg.Core.CleanerInterval, manager.CleanExpired)
	cleaner.Start()
	slog.Info("cleaner started")
	defer cleaner.Stop()

	srv := server.NewServer(&cfg.Server, manager)

	serverErrChan := make(chan error, 1)
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, net.ErrClosed) {
			serverErrChan <- err
		}
		close(serverErrChan)
	}()

	select {
	case <-ctx.Done():
		slog.Info("received shutdown signal, stopping application...")
	case err := <-serverErrChan:
		if err != nil {
			slog.Error("server failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("failed to shutdown server gracefully", slog.String("error", err.Error()))
	} else {
		slog.Info("server stopped gracefully")
	}
}
