package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedisConn(cfg Config) (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		},
	)
	if cmd := client.Ping(); cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return client, nil
}
