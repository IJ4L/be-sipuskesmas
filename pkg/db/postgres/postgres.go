package postgres

import (
	"context"
	"time"

	"github.com/IJ4L/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(cfg utils.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	DBUrl := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
	poolConfig, err := pgxpool.ParseConfig(DBUrl)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = 20
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 15 * time.Minute

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
