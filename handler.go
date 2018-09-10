package pgbatch

import (
	"time"

	"database/sql"

	pgservice "github.com/behrang/go-pgservice"
	_ "github.com/lib/pq" // Just load pq library.
)

// PostgresHandler contains connection pool.
type PostgresHandler struct {
	Pool *sql.DB
}

// Close handler connection pool.
func (handler *PostgresHandler) Close() error {
	return handler.Pool.Close()
}

// New connects to DB and returns wrapped pool.
func New(service, file string) (*PostgresHandler, error) {

	// Read service section of config file.
	err := pgservice.Apply(service, file)
	if err != nil {
		return nil, err
	}

	// Connect to db.
	pool, err := sql.Open("postgres", "")
	if err != nil {
		return nil, err
	}

	// Ping db.
	err = pool.Ping()
	if err != nil {
		return nil, err
	}

	pool.SetMaxOpenConns(1)
	pool.SetMaxIdleConns(1)
	pool.SetConnMaxLifetime(60 * time.Second)

	return &PostgresHandler{pool}, nil
}
