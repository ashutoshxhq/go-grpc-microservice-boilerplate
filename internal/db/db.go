package db

import (
	"context"
	"fmt"
	"time"

	"github.com/egnitelabs/engine/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, c *config.Config) (*sqlx.DB, context.CancelFunc, error) {

	var dsn string
	dsn = fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s connect_timeout=%d sslmode=disable",
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDatabase,
		10,
	)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, cancel, err
	}

	db.SetMaxIdleConns(c.PostgresMaxConnections)
	db.SetMaxIdleConns(c.PostgresMaxConnections * 10)
	db.SetConnMaxLifetime(time.Minute * time.Duration(c.PostgresMaxConnectionLifetime))

	return db, cancel, nil
}
