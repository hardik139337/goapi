package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient = redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_URL")})
