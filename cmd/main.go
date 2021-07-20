package main

import (
	"flag"
	"log"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/server"
	_ "github.com/lib/pq"
)

var (
	configPath string
)

func init() {
	flag.StringVar(
		&configPath,
		"config-path",
		"./config/config.yml",
		"path to configuration file",
	)
}

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
