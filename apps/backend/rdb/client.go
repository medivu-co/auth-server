package rdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"medivu.co/auth/envs"
)

type redisClient struct {
	*redis.Client
	connected bool
}

var client *redisClient = &redisClient{
	Client:    nil,
	connected: false,
}

func Connect() error {
	client.Client = redis.NewClient(&redis.Options{
		Addr:     envs.RedisAddr(),
		Password: envs.RedisPassword(),
		DB:       envs.RedisDB(),
	})
	// Try to ping Redis to check connection
	if err := client.Client.Ping(context.Background()).Err(); err != nil {
		client.connected = false
		panic(errors.New("failed to ping to Redis: " + err.Error()))
	}
	fmt.Println("Connected to Redis")
	client.connected = true
	return nil
}

func Set(key string, value string, expiration time.Duration) error {
	if !client.connected {
		return errors.New("redis client is not connected")
	}
	return client.Set(context.Background(), key, value, expiration).Err()
}
func Get(key string) (string, error) {
	if !client.connected {
		return "", errors.New("redis client is not connected")
	}
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func HSet(key string, value IRedisHash, expiration time.Duration) error {
	if !client.connected {
		return errors.New("redis client is not connected")
	}
	if err := client.HSet(context.Background(), key, value.BuildHash()).Err(); err != nil {
		return err
	}
	// set Expiration
	return client.Expire(context.Background(), key, expiration).Err()
}

func HGetAll(key string) (map[string]string, error) {
	if !client.connected {
		return nil, errors.New("redis client is not connected")
	}
	val, err := client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func RPush(key string, value string) error {
	if !client.connected {
		return errors.New("redis client is not connected")
	}
	return client.RPush(context.Background(), key, value).Err()
}

func LPop(key string) (*string, error) {
	if !client.connected {
		return nil, errors.New("redis client is not connected")
	}
	val, err := client.LPop(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &val, nil
}

type IRedisHash interface {
	BuildHash() map[string]string
}
