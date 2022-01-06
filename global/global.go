package global

import (
	"crypto/rsa"
	"github.com/go-redis/redis"
)

var (
	CONFIG *Config
	RDB *redis.Client
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
)
