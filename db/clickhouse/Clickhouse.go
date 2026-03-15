package clickhouse

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func NewClickhouseConn(cfg Config) (driver.Conn, error) {
	conn, er := clickhouse.Open(
		&clickhouse.Options{
			Addr: []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
			Auth: clickhouse.Auth{
				Database: cfg.Database,
				Username: cfg.User,
				Password: cfg.Pass,
			},
		},
	)
	if er != nil {
		return nil, er
	}
	if er := conn.Ping(context.Background()); er != nil {
		return nil, er
	}
	return conn, nil
}
