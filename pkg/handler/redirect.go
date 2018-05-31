package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/seregayoga/shorturl/pkg/url"
)

// Redirect handler for redirects
type Redirect struct {
	shortener *url.Shortener
}

// NewRedirect creates new handler
func NewRedirect(shortener *url.Shortener) *Redirect {
	return &Redirect{
		shortener,
	}
}

// Handler handler func for router
func (h *Redirect) Handler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	longURL, err := h.shortener.GetLongURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
