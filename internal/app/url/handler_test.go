package url

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erickir/tinyurl/internal/app/url/models"
	"github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	c := require.New(t)

	service := URLService{}

	handler := New(service)
	c.NotNil(handler)
}

func TestRoutes(t *testing.T) {
	c := require.New(t)

	service := URLService{}

	handler := New(service)
	c.NotNil(handler)

	routes := handler.Routes()
	c.NotNil(routes)
}

func TestGetLongUrlHandlerSuccess(t *testing.T) {
	c := require.New(t)

	mockStorage := storage.Mock()

	tinyURL := &models.TinyURL{
		ShortID: "xyz",
		LongURL: "xyz.com",
	}

	mockStorage.SetMockTinyURL(tinyURL)

	urlService := NewURLService(mockStorage)
	handler := New(urlService).getLongUrl()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/xyz", nil)

	handler.ServeHTTP(rec, req)

	c.Equal(http.StatusTemporaryRedirect, rec.Result().StatusCode)
	c.Equal(rec.Result().Header.Get("Location"), tinyURL.LongURL)

}
