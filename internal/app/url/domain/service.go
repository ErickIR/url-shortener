package domain

import (
	"context"
	"net/url"

	"github.com/erickir/tinyurl/internal/app/url/models"
	urlStorage "github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/erickir/tinyurl/pkg/base62"
)

type Service interface {
	GetLongURL(ctx context.Context, shortID string) (string, error)
	SaveURL(ctx context.Context, rawURL string) (*models.TinyURLResponse, error)
}

type URLService struct {
	storage urlStorage.Storage
}

func NewURLService(storage urlStorage.Storage) *URLService {
	return &URLService{
		storage: storage,
	}
}

func (service *URLService) GetLongURL(ctx context.Context, shortID string) (string, error) {
	tinyUrl, err := service.storage.GetTinyURLByID(ctx, shortID)
	if err != nil {
		return "", err
	}

	return tinyUrl.LongURL, nil
}

func (service *URLService) SaveURL(ctx context.Context, rawURL string) (*models.TinyURLResponse, error) {
	_, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	shortURL := base62.StringToBase64(rawURL)

	tinyURL := &models.TinyURL{
		ShortID: shortURL,
		LongURL: rawURL,
	}

	if err := service.storage.SaveURL(ctx, tinyURL); err != nil {
		return nil, err
	}

	return tinyURL.ToResponse(), nil
}
