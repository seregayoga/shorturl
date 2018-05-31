package handler

import (
	"encoding/json"
	"net/http"

	"github.com/seregayoga/shorturl/pkg/config"
	"github.com/seregayoga/shorturl/pkg/url"
)

// CreateRequest create request with long url
type CreateRequest struct {
	LongURL string `json:"long_url"`
}

// CreateResponse create response with short url
type CreateResponse struct {
	ShortURL string `json:"short_url"`
}

// Create handler for creating short urls
type Create struct {
	shortener *url.Shortener
	cfg       *config.Config
}

// NewCreate creates new handler
func NewCreate(shortener *url.Shortener, cfg *config.Config) *Create {
	return &Create{
		shortener: shortener,
		cfg:       cfg,
	}
}

// Handler handler func for router
func (h *Create) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createRequest CreateRequest
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil || createRequest.LongURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL, err := h.shortener.CreateShortURL(createRequest.LongURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createResponse := CreateResponse{
		ShortURL: h.cfg.RedirectHost + shortURL,
	}
	json.NewEncoder(w).Encode(createResponse)
}
