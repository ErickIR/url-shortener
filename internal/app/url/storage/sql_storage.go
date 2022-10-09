package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/erickir/tinyurl/internal/app/url/models"
	"github.com/erickir/tinyurl/pkg/mssql"
)

var (
	// ErrShortURLNotFound
	ErrShortURLNotFound = errors.New("short url not found")

	tinyUrlsTableName = "TINY_URLS"
)

type URLStorage struct {
	db mssql.SQL
}

func NewURLStorage(db mssql.SQL) *URLStorage {
	return &URLStorage{
		db: db,
	}
}

func (storage *URLStorage) GetTinyURLByID(ctx context.Context, shortID string) (*models.TinyURL, error) {
	query := fmt.Sprintf("SELECT short_id, long_url FROM %s WHERE short_id = @p1", tinyUrlsTableName)

	rows, err := storage.db.QueryContext(ctx, query, shortID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			panic(err)
		}
	}()

	tinyURL := &models.TinyURL{}
	if !rows.Next() {
		return nil, ErrShortURLNotFound
	}

	if err := rows.Scan(&tinyURL.ShortID, &tinyURL.LongURL); err != nil {
		return nil, err
	}

	return tinyURL, nil
}

func (storage *URLStorage) SaveURL(ctx context.Context, tinyURL *models.TinyURL) error {
	query := fmt.Sprintf("INSERT INTO %s (short_id, long_url) VALUES (@p1, @p2)", tinyUrlsTableName)

	_, err := storage.db.ExecContext(ctx, query, tinyURL.ShortID, tinyURL.LongURL)
	return err
}
