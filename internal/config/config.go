package config

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"sURL/internal/api"
	"sURL/internal/storage"
)

type Config struct {
	Storage          storage.Repository
	BasicResaultAddr string
	Address          string
}

func (c *Config) Run() {
	fmt.Printf("Server address: %s\n\rBasic resault address: %s\n\r", c.Address, c.BasicResaultAddr)
	if err := http.ListenAndServe(c.Address, api.NewRouter(c.Storage, c.BasicResaultAddr)); err != nil {
		slog.Error(err.Error())
	}
}

func Configure(storage storage.Repository) *Config {
	config := &Config{
		Storage:          storage,
		BasicResaultAddr: "",
		Address:          "",
	}

	flag.StringVar(&config.BasicResaultAddr, "b", "http://localhost:8080", "The server basic redirect url")
	flag.StringVar(&config.Address, "a", ":8080", "The server Address and port")
	flag.Parse()

	return config
}
