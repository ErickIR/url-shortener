package url

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/erickir/tinyurl/pkg/api"
	"github.com/gofiber/fiber/v2"
)

const (
	shortIDParamKey    = "SHORT_URL_ID"
	urlResourcePathKey = "/url"
)

type TinyUrlHandler struct {
	service Service
}

type saveURLRequest struct {
	LongURL string `json:"long_url,omitempty"`
}

func New(service Service) *TinyUrlHandler {
	return &TinyUrlHandler{
		service: service,
	}
}

func (h TinyUrlHandler) Path() string {
	return urlResourcePathKey
}

func (h TinyUrlHandler) Routes(app fiber.Router) {
	app.Get("/:"+shortIDParamKey, h.GetLongUrl())
	app.Post("/", h.ShortenUrl())
}

func (h *TinyUrlHandler) GetLongUrl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		shortID := c.Params(shortIDParamKey, "")

		longUrl, err := h.service.GetLongURL(ctx, shortID)
		if errors.Is(err, ErrTinyURLNotFound) {
			return api.RespondError(c, http.StatusNotFound, api.ResourceNotFoundError)
		}

		if err != nil {
			return api.RespondError(c, http.StatusInternalServerError, api.InternalServerError)
		}

		return c.Redirect(longUrl)
	}
}

func (h *TinyUrlHandler) ShortenUrl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		var requestBody saveURLRequest

		if err := json.Unmarshal(c.Body(), &requestBody); err != nil {
			return api.RespondError(c, http.StatusNotFound, api.InternalServerError)
		}

		tinyURL, err := h.service.SaveURL(ctx, requestBody.LongURL)
		if errors.Is(err, ErrInvalidURLReceived) {
			return api.RespondError(c, http.StatusBadRequest, api.NewErrorResponse(err))
		}

		if err != nil {
			return api.RespondError(c, http.StatusInternalServerError, api.InternalServerError)
		}

		return api.RespondWithJSON(c, http.StatusOK, tinyURL.ToResponse())
	}
}
