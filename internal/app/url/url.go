package url

import (
	"github.com/erickir/tinyurl/internal/app/url/domain"
	urlHandlers "github.com/erickir/tinyurl/internal/app/url/handlers"
	"github.com/erickir/tinyurl/internal/app/url/storage"
)

func Setup() *urlHandlers.Handler {
	inMemoryDB := storage.NewInMemoryDB()

	urlService := domain.NewURLService(inMemoryDB)

	return urlHandlers.New(urlService)
}
