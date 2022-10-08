package url

import (
	"context"
	"testing"

	"github.com/erickir/tinyurl/internal/app/url/models"
	"github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/stretchr/testify/require"
)

func TestNewURLService(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	service := NewURLService(mockStorage)
	c.NotNil(service)
}

func TestGetLongURLSuccess(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	service := NewURLService(mockStorage)

	longURL, err := service.GetLongURL(context.Background(), "short-id")
	c.NoError(err)
	c.Equal(storage.MockTinyURL.LongURL, longURL)
}

func TestGetLongURLForcedFailure(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	mockStorage.ForceFailure = true
	defer func() {
		mockStorage.ForceFailure = false
	}()

	service := NewURLService(mockStorage)

	_, err := service.GetLongURL(context.Background(), "short-id")
	c.ErrorIs(err, storage.ErrForcedFailure)
}

func TestGetLongURLNotFoundError(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	mockStorage.ForceNotFound = true
	defer func() {
		mockStorage.ForceNotFound = false
	}()

	service := NewURLService(mockStorage)

	_, err := service.GetLongURL(context.Background(), "short-id")
	c.ErrorIs(err, ErrTinyURLNotFound)
}

func TestSaveURLSuccess(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	service := NewURLService(mockStorage)

	response, err := service.SaveURL(context.Background(), "wwww.google.com")
	c.NoError(err)
	c.IsType(&models.TinyURLResponse{}, response)
}

func TestSaveURLWithInvalidURL(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	service := NewURLService(mockStorage)

	deleteCtrlChar := string(rune(0x7f))

	_, err := service.SaveURL(context.Background(), deleteCtrlChar)
	c.ErrorIs(err, ErrInvalidURLReceived)
}

func TestSaveURLWithForcedFailure(t *testing.T) {
	c := require.New(t)

	mockStorage := &storage.StorageMock{}

	mockStorage.ForceFailure = true
	defer func() {
		mockStorage.ForceFailure = false
	}()

	service := NewURLService(mockStorage)

	_, err := service.SaveURL(context.Background(), "wwww.google.com")
	c.ErrorIs(err, storage.ErrForcedFailure)
}
