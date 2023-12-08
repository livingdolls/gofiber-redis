package database

import (
	"fmt"
	"gofiber-redis/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connected to redis database")

	return rdb
}