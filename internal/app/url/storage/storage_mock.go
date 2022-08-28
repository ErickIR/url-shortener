package storage

import (
	"context"
	"errors"

	"github.com/erickir/tinyurl/internal/app/url/models"
)

var (
	ErrForcedFailure = errors.New("forced failure")

	MockTinyURL = &models.TinyURL{
		ShortID: "short-id",
		LongURL: "www.google.com",
	}
)

type StorageMock struct {
	ForceNotFound bool
	ForceFailure  bool
}

func (s StorageMock) GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error) {
	if s.ForceFailure {
		return nil, ErrForcedFailure
	}

	if s.ForceNotFound {
		return nil, ErrShortURLNotFound
	}

	return MockTinyURL, nil
}

func (s StorageMock) SaveURL(ctx context.Context, tinyURL *models.TinyURL) error {
	if s.ForceFailure {
		return ErrForcedFailure
	}

	return nil
}
