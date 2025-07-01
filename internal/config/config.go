package config

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sURL/internal/api"
	"sURL/internal/storage"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Storage          storage.Repository
	BasicResaultAddr string `env:"BASE_URL"`
	Address          string `env:"SERVER_ADDRESS"`
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

	if err := env.Parse(config); err != nil {
		log.Fatal(err.Error())
	}

	return config
}

func (c *Config) Run() {
	fmt.Printf("Server address: %s\n\rBasic resault address: %s\n\r", c.Address, c.BasicResaultAddr)
	if err := http.ListenAndServe(c.Address, api.NewRouter(c.Storage, c.BasicResaultAddr)); err != nil {
		log.Fatal(err.Error())
		return
	}
}
