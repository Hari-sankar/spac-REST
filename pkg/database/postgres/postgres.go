package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxOpenConns    = 60  // Equivalent to MaxConns in pgxpool
	connMaxLifetime = 120 // This is unused in pgxpool but can be managed manually
	maxIdleConns    = 30  // Not directly applicable in pgxpool
	connMaxIdleTime = 20  // pgxpool config sets max idle time
)

// NewPsqlDB initializes and returns a pgxpool.Pool connection pool
func NewPsqlDB(connectionURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	config.MaxConns = int32(maxOpenConns)
	config.MaxConnLifetime = time.Duration(connMaxLifetime) * time.Second
	config.MaxConnIdleTime = time.Duration(connMaxIdleTime) * time.Second

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Test the connection with Ping
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
