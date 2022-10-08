package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/erickir/tinyurl/internal/app/url/domain"
	"github.com/erickir/tinyurl/pkg/api"
	"github.com/go-chi/chi/v5"
)

const (
	shortIDParamKey    = "SHORT_URL_ID"
	urlResourcePathKey = "/url"
)

type TinyUrlHandler struct {
	service domain.Service
}

type saveURLRequest struct {
	LongURL string `json:"long_url,omitempty"`
}

func New(service domain.Service) *TinyUrlHandler {
	return &TinyUrlHandler{
		service: service,
	}
}

func (h TinyUrlHandler) Path() string {
	return urlResourcePathKey
}

func (h TinyUrlHandler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get(fmt.Sprintf("/{%s}", shortIDParamKey), h.getLongUrl())
	r.Post("/", h.shortenUrl())

	return r
}

func (h *TinyUrlHandler) getLongUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		shortID := chi.URLParam(r, shortIDParamKey)

		longUrl, err := h.service.GetLongURL(ctx, shortID)
		if errors.Is(err, domain.ErrTinyURLNotFound) {
			api.RespondWithJSON(w, http.StatusNotFound, api.ResourceNotFoundError)
			return
		}

		if err != nil {
			api.RespondWithJSON(w, http.StatusInternalServerError, api.InternalServerError)
			return
		}

		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	}
}

func (h *TinyUrlHandler) shortenUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var requestBody saveURLRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			api.RespondWithJSON(w, http.StatusInternalServerError, api.InternalServerError)
			return
		}

		tinyURL, err := h.service.SaveURL(ctx, requestBody.LongURL)
		if errors.Is(err, domain.ErrInvalidURLReceived) {
			api.RespondWithJSON(w, http.StatusBadRequest, api.NewErrorResponse(err))
			return
		}

		if err != nil {
			api.RespondWithJSON(w, http.StatusInternalServerError, api.InternalServerError)
			return
		}

		api.RespondWithJSON(w, http.StatusOK, tinyURL.ToResponse())
	}
}
