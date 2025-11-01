package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis() error {
	addr := os.Getenv("REDIS_ADDR")
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Username: username,
			Password: password,
			DB:       0,
		},
	)

	RDB = client

	if err := client.Set(ctx, "connected", "Yes Connected", 0).Err(); err != nil {
		return err
	}
	result, err := client.Get(ctx, "connected").Result()
	if err != nil {
		return err
	}
	println(result)
	return nil
}
