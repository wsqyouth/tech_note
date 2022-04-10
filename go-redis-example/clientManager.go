package main

import (
	redis "github.com/go-redis/redis"
)

// NewRedisClient create redis client instance
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
