package config

import (
	"flag"
	"log/slog"
	"net/http"
	"sURL/internal/api"
	"sURL/internal/storage"
)

type Config struct {
	storage          storage.Repository
	basicResaultAddr string
	address          string
}

func (c *Config) Run() {
	if err := http.ListenAndServe(c.address, api.NewRouter(c.storage)); err != nil {
		slog.Error(err.Error())
	}
}

func Configure(storage storage.Repository) Config {
	var config Config
	config.storage = storage
	flag.StringVar(&config.basicResaultAddr, "b", "http://localhost:8080", "The server basic redirect url")
	flag.StringVar(&config.address, "a", ":8080", "The server address and port")
	flag.Parse()
	return config
}
