package main

import (
	"sURL/internal/config"
	"sURL/internal/storage/memory"
)

func main() {
	storage := memory.MemStorage{}

	app := config.Configure(&storage)
	app.Run()
}
