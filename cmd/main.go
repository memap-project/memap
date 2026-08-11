package main

import (
	"context"
	"log"

	"github.com/memap-project/memap/config"
	"github.com/memap-project/memap/core/clean"
	"github.com/memap-project/memap/core/ns"
	"github.com/memap-project/memap/server"
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
