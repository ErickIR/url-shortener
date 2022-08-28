package mssql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	connectionStringFormat = "server=%s;user id=%s;password=%s;port=%s;database=%s;"
)

var (
	// SQLOpen open function for mock
	SQLOpen = sql.Open
)

type SQL interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PingContext(ctx context.Context) error
}

type SQLClient struct {
	db *sql.DB
}

func getConnectionString(server, userId, password, port, database string) string {
	return fmt.Sprintf(connectionStringFormat, server, userId, password, port, database)
}

func NewClient(ctx context.Context, server, userId, password, port, database string) (*SQLClient, error) {
	connectionString := getConnectionString(server, userId, password, port, database)

	conn, err := SQLOpen("sqlserver", connectionString)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database failed: %w", err)
	}

	return &SQLClient{
		db: conn,
	}, nil
}

func (c *SQLClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *SQLClient) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}

func (c *SQLClient) PingContext(ctx context.Context) error {
	return c.db.PingContext(ctx)
}
