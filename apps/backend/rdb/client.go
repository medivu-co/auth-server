package rdb

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"medivu.co/auth/envs"
	"medivu.co/auth/logger"
)

var Client *redis.Client	

func Connect() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     envs.RedisAddr(),
		Password: envs.RedisPassword(),
		DB:       envs.RedisDB(),
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})
	// Try to ping Redis to check connection
	if err := Client.Ping(context.Background()).Err(); err != nil {
		logger.Get().Fatal("failed to ping to Redis: " + err.Error())
		return err
	}
	logger.Get().Info("Connected to Redis")
	return nil
}

func Close() error {
	if err := Client.Close(); err != nil {
		logger.Get().Fatal("failed to close Redis connection: " + err.Error())
		return err
	}
	logger.Get().Info("Redis connection closed")
	return nil
}