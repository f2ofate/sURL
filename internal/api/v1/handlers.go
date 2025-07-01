package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"sURL/internal/storage"

	"github.com/go-chi/chi/v5"
)

type requset struct {
	URL string `json:"url"`
}

type response struct {
	Result string `json:"result"`
}

// StoreURL получает тело запроса и сохраняет его в хранилище
func StoreURL(s storage.Repository, resaultAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requset

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err = json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
		}

		shortUrl, ok := s.AddShortURL(req.URL)
		if !ok {
			http.Error(w, "Invalid URL format. Must be http://... or https://...", http.StatusBadRequest)
			return
		}

		res := response{Result: resaultAddr + "/" + shortUrl}

		jsonRes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Error marshalling response body: "+err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonRes)
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
