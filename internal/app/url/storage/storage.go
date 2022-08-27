package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/erickir/tinyurl/internal/app/url/models"
)

var (
	ErrTinyURLNotFound = errors.New("error tiny url not found")
)

type Storage interface {
	GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error)
	SaveURL(ctx context.Context, tinyURL *models.TinyURL) error
}

type InMemory struct {
	tinyURLs []*models.TinyURL
	mutex    sync.Mutex
}

func NewInMemoryDB() *InMemory {
	return &InMemory{
		tinyURLs: make([]*models.TinyURL, 0),
		mutex:    sync.Mutex{},
	}
}

func (db *InMemory) GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	for _, url := range db.tinyURLs {
		if shortID == url.ShortID {
			return url, nil
		}
	}

	return nil, ErrTinyURLNotFound
}

func (db *InMemory) SaveURL(ctx context.Context, tinyURL *models.TinyURL) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.tinyURLs = append(db.tinyURLs, tinyURL)
	return nil
}
