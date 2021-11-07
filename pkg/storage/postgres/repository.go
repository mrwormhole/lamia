package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MrWormHole/perseusCMS/pkg/storage/postgres/sqlc"

	"github.com/MrWormHole/perseusCMS/pkg/storage"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type Storage struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewStorage(cfg *storage.Config) (*Storage, error) {
	if cfg == nil || cfg.ConnectionString() == "" {
		return nil, fmt.Errorf("invalid configuration for the given connection string: %s", cfg.ConnectionString())
	}

	if cfg.DriverType != storage.PostgresSQL {
		return nil, fmt.Errorf("invalid storage type for the given driver type: %s", cfg.DriverType)
	}

	db, err := sql.Open(driverName, cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %s", err)
	}
	queries := sqlc.New(db)

	return &Storage{
		db:      db,
		queries: queries,
	}, nil
}
