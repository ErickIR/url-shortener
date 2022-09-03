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

// Mock returns a sql mock client for tests
func Mock() *SQLMockClient {
	return &SQLMockClient{}
}

// ActivateForcedFailure activates the force failure of functions
func (c *SQLMockClient) ActivateForcedFailure() {
	c.forceFailure = true
}

// DeactivateForcedFailure deactivates the force failure of functions
func (c *SQLMockClient) DeactivateForcedFailure() {
	c.forceFailure = false
}

// ExecContext
func (c *SQLMockClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if c.forceFailure {
		return nil, ErrForcedFailure
	}

	return &MockResult{}, nil
}

// QueryContext
func (c *SQLMockClient) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if c.forceFailure {
		return nil, ErrForcedFailure
	}

	return &sql.Rows{}, nil
}

// PingContext
func (c *SQLMockClient) PingContext(ctx context.Context) error {
	if c.forceFailure {
		return ErrForcedFailure
	}

	return nil
}
