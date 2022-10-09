package storage

import (
	"testing"

	"github.com/erickir/tinyurl/pkg/mssql"
	"github.com/stretchr/testify/require"
)

var (
	_ Storage = &URLStorage{}
)

func TestNewSQLStorage(t *testing.T) {
	c := require.New(t)

	mockClient := mssql.SQLMockClient{}

	sqlStorage := NewURLStorage(&mockClient)

	c.NotNil(sqlStorage)
}
