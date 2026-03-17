package redis

import (
	"context"

	"blog.kimimao.com/backend/config"
	"github.com/redis/go-redis/v9"
)

// Client Redis客户端
var Client *redis.Client

// InitRedis 初始化Redis连接
func InitRedis(cfg *config.RedisConfig) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return Client.Ping(context.Background()).Err()
}
