package app

import (
	"log/slog"
	"net/http"
	"sURL/internal/api"
	"sURL/internal/storage/memory"
)

func Run(port string) {
	storage := memory.MemStorage{}

	if err := http.ListenAndServe(":"+port, api.NewRouter(&storage)); err != nil {
		slog.Error(err.Error())
	}
}
