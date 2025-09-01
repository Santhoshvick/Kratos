package biz

import (
	"github.com/redis/go-redis/v9"
)



func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // correct field name is Addr
		Password: "",               // no password set
		DB:       0,                // default DB
	})
}
