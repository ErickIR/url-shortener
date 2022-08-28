package url

import (
	"github.com/erickir/tinyurl/internal/app/url/domain"
	urlHandlers "github.com/erickir/tinyurl/internal/app/url/handlers"
	"github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/erickir/tinyurl/pkg/mssql"
)

func Setup(db mssql.SQL) *urlHandlers.Handler {
	store := storage.NewSQLStorage(db)

	urlService := domain.NewURLService(store)

	return urlHandlers.New(urlService)
}
