package v1

import (
	"fmt"
	"io"
	"net/http"
	"sURL/internal/storage"

	"github.com/go-chi/chi/v5"
)

// StoreURL получает тело запроса и сохраняет его в хранилище
func StoreURL(s storage.Repository, resaultAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		shortUrl, ok := s.AddShortURL(string(body))
		if !ok {
			http.Error(w, "Invalid URL format. Must be http://... or https://...", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("%s/%s", resaultAddr, shortUrl)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(response))
	}
}

// RedirectURL при переходе на сокращённую ссылку перенаправляет на оригинальную
func RedirectURL(s storage.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		originUrl := s.GetOriginURL(id)

		w.Header().Set("Content-Type", "text/plain")
		http.Redirect(w, r, originUrl, http.StatusTemporaryRedirect)
	}
}
