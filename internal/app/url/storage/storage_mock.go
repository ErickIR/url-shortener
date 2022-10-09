package storage

import (
	"context"
	"errors"

	"github.com/erickir/tinyurl/internal/app/url/models"
)

var (
	ErrForcedFailure = errors.New("forced failure")
)

type StorageMock struct {
	ForceNotFound bool
	ForceFailure  bool
	mockTinyURLs  []*models.TinyURL
}

func Mock() *StorageMock {
	return &StorageMock{}
}

func (s *StorageMock) SetMockTinyURL(tinyURL *models.TinyURL) {
	s.mockTinyURLs = append(s.mockTinyURLs, tinyURL)
}

func (s StorageMock) GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error) {
	if s.ForceFailure {
		return nil, ErrForcedFailure
	}

	if s.ForceNotFound {
		return nil, ErrShortURLNotFound
	}

	return lookupShortURL(s.mockTinyURLs, shortID)
}

func lookupShortURL(mockURLs []*models.TinyURL, shortID string) (*models.TinyURL, error) {
	for _, url := range mockURLs {
		if url.ShortID == shortID {
			return url, nil
		}
	}

	return nil, ErrShortURLNotFound
}

func (s StorageMock) SaveURL(ctx context.Context, tinyURL *models.TinyURL) error {
	if s.ForceFailure {
		return ErrForcedFailure
	}

	s.SetMockTinyURL(tinyURL)
	return nil
}
