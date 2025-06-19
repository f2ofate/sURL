package app

import (
	"log/slog"
	"net/http"
	shttp "sURL/internal/controller/http"
	"sURL/internal/storage/memory"
)

func Run(port string) {
	storage := memory.MemStorage{}

	if err := http.ListenAndServe(":"+port, shttp.NewRouter(&storage)); err != nil {
		slog.Error(err.Error())
	}
}
