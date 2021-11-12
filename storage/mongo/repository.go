package mongo

import (
	"context"
	"fmt"
	"github.com/MrWormHole/perseusCMS/storage"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

type Storage struct {
	db *mongo.Database
}

func NewStorage(cfg *storage.Config) (*Storage, error) {
	if cfg == nil || cfg.ConnectionString() == "" {
		return nil, fmt.Errorf("invalid configuration for the given connection string: %s", cfg.ConnectionString())
	}

	if cfg.DriverType != storage.MongoDB {
		return nil, fmt.Errorf("invalid storage type for the given driver type: %s", cfg.DriverType)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.ConnectionString()))
	if err != nil {
		return nil, fmt.Errorf("failed to init client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("failed to open db connection: %v", err)
	}
	db := client.Database(cfg.DatabaseName())

	return &Storage{
		db: db,
	}, nil
}
