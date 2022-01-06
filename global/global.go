package global

import (
	"crypto/rsa"
	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
	PrivateKey *rsa.PrivateKey
)
