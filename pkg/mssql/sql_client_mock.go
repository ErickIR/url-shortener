package mssql

import (
	"context"
	"database/sql"
	"errors"
)

var (
	// ErrForcedFailure
	ErrForcedFailure = errors.New("forced failure")
)

type SQLMockClient struct {
	forceFailure bool
}

func (c *SQLMockClient) ActivateForcedFailure() {
	c.forceFailure = true
}

func (c *SQLMockClient) DeactivateForcedFailure() {
	c.forceFailure = false
}

func (c *SQLMockClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if c.forceFailure {
		return nil, ErrForcedFailure
	}

	return &MockResult{}, nil
}

func (c *SQLMockClient) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if c.forceFailure {
		return nil, ErrForcedFailure
	}

	return &sql.Rows{}, nil
}

func (c *SQLMockClient) PingContext(ctx context.Context) error {
	if c.forceFailure {
		return ErrForcedFailure
	}

	return nil
}
