package redisClient

import (
	"github.com/redis/go-redis/v9"
)

func New(address, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return client
}
