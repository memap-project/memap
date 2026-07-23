package main

import (
	"log"

	"github.com/dmi3midd/memap/config"
	"github.com/dmi3midd/memap/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	server := server.NewServer(&cfg.Server)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
