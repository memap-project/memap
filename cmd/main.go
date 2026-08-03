package main

import (
	"context"
	"log"

	"github.com/dmi3midd/memap/config"
	"github.com/dmi3midd/memap/core/clean"
	"github.com/dmi3midd/memap/core/ns"
	"github.com/dmi3midd/memap/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx := context.Background()
	manager := ns.NewNamespaceManager()
	cleaner := clean.NewCleaner(ctx, cfg.Core.CleanerInterval, manager.Clean)
	cleaner.Start()
	defer cleaner.Stop()

	server := server.NewServer(&cfg.Server, manager)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
