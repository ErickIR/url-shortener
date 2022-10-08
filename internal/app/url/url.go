package url

import (
	"github.com/erickir/tinyurl/internal/app/url/storage"
	"github.com/erickir/tinyurl/pkg/mssql"
)

func Setup(db mssql.SQL) *TinyUrlHandler {
	store := storage.NewSQLStorage(db)

	urlService := NewURLService(store)

	return New(urlService)
}
