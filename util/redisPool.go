package util

import (
	"github.com/garyburd/redigo/redis"
)

var (
	RedisClient *redis.Pool
)
