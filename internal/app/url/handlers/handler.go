package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/erickir/tinyurl/internal/app/url/domain"
	"github.com/go-chi/chi/v5"
)

const (
	shortIDParamKey = "SHORT_URL_ID"
)

type Handler struct {
	service domain.Service
}

type saveURLRequest struct {
	LongURL string `json:"long_url,omitempty"`
}

func New(service domain.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get(fmt.Sprintf("/{%s}", shortIDParamKey), h.getLongUrl())
	r.Post("/", h.shortenUrl())

	return r
}

func (h *Handler) getLongUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		shortID := chi.URLParam(r, shortIDParamKey)

		longUrl, err := h.service.GetLongURL(ctx, shortID)
		if err != nil {
			log.Fatal("ERROR IN API")
		}

		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	}
}

func (h *Handler) shortenUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var requestBody saveURLRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			log.Fatal("ERROR IN API")
		}

		response, err := h.service.SaveURL(ctx, requestBody.LongURL)
		if err != nil {
			log.Fatal("ERROR IN API")
		}

		json.NewEncoder(w).Encode(response)
	}
}
