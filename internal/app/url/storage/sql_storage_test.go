package storage

import (
	"testing"

	"github.com/erickir/tinyurl/pkg/mssql"
	"github.com/stretchr/testify/require"
)

var (
	_ Storage = &SQLStorage{}
)

func TestNewSQLStorage(t *testing.T) {
	c := require.New(t)

	mockClient := mssql.SQLMockClient{}

	sqlStorage := NewSQLStorage(&mockClient)

	c.NotNil(sqlStorage)
}
