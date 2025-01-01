package redis

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/pkg/config"
	redis "github.com/redis/go-redis/v9"
)

func New(cfg *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
