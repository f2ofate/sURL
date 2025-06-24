package api

import (
	"net/http"
	"sURL/internal/api/v1"
	"sURL/internal/storage"

	"github.com/go-chi/chi/v5"
)

func NewRouter(s storage.Repository) http.Handler {
	r := chi.NewRouter()

	r.Post("/", v1.StoreURL(s))
	r.Get("/{id}", v1.RedirectURL(s))

	return r
}
