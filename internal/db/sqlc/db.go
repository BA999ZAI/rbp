package sqlc

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"rbp/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

func NewDB(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.DB.User, config.DB.Password),
		Host:   fmt.Sprintf("%s:%d", config.DB.Host, config.DB.Port),
		Path:   config.DB.DBName,
	}

	q := make(url.Values)
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()

	poolConfig, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: tracelog.LoggerFunc(
			func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
				args := make([]any, 0, len(data))
				for k, v := range data {
					args = append(args, slog.Any(k, v))
				}
			},
		),
		LogLevel: tracelog.LogLevelTrace,
	}

	// TODO: Add and use input context?
	c, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := c.Ping(ctx); err != nil {
		return nil, err
	}

	return c, nil
}
