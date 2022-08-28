package storage

import (
	"context"

	"github.com/erickir/tinyurl/internal/app/url/models"
)

type Storage interface {
	GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error)
	SaveURL(ctx context.Context, tinyURL *models.TinyURL) error
}
