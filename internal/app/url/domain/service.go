package domain

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/erickir/tinyurl/internal/app/url/models"
	"github.com/erickir/tinyurl/internal/app/url/storage"
	urlStorage "github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/erickir/tinyurl/pkg/base62"
)

var (
	ErrTinyURLNotFound = errors.New("short url not found")

	ErrInvalidURLReceived = errors.New("invalid url received")
)

type Service interface {
	GetLongURL(ctx context.Context, shortID string) (string, error)
	SaveURL(ctx context.Context, rawURL string) (*models.TinyURL, error)
}

type URLService struct {
	storage urlStorage.Storage
}

func NewURLService(storage urlStorage.Storage) *URLService {
	return &URLService{
		storage: storage,
	}
}

func (service URLService) GetLongURL(ctx context.Context, shortID string) (string, error) {
	tinyUrl, err := service.storage.GetTinyURLByID(ctx, shortID)
	if errors.Is(err, storage.ErrShortURLNotFound) {
		return "", ErrTinyURLNotFound
	}

	if err != nil {
		return "", err
	}

	return tinyUrl.LongURL, nil
}

func (service URLService) SaveURL(ctx context.Context, rawURL string) (*models.TinyURL, error) {
	_, err := url.Parse(rawURL)
	if err != nil {
		return nil, ErrInvalidURLReceived
	}

	shortURL := base62.StringToBase64(rawURL)

	tinyURL := &models.TinyURL{
		ShortID: shortURL,
		LongURL: rawURL,
	}

	err = service.storage.SaveURL(ctx, tinyURL)
	if err != nil {
		fmt.Println("ERROR SAVING URL: ", err.Error())
		return nil, err
	}

	return tinyURL, nil
}
