package db

import (
	"RushBananaBet/pkg/logger"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	// postgres://user:pass@localhost:5432/dbname?sslmode=disable

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.Fatal("Error parse config for pgxpool", "db - NewPostgresPool()", err)
		return nil, err
	}

	logger.Debug("Success parse DB Config", "db-NewPostgresPool()", nil)

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Fatal("Error create pgxpool", "db - NewPostgresPool()", err)
		return nil, err
	}

	logger.Debug("Success create pgxpool", "db-NewPostgresPool()", nil)

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal("Error database is not responding", "db - NewPostgresPool()", err)
		return nil, err
	}

	logger.Debug("Success ping databse", "db-NewPostgresPool()", nil)

	return pool, nil
}
