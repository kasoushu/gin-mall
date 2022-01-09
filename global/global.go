package global

import (
	"crypto/rsa"
	"database/sql"
	"gin_mall/global/config"
	"github.com/go-redis/redis"
)

var (
	Config     *config.Config
	RDB        *redis.Client
	MDB        *sql.DB
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)
