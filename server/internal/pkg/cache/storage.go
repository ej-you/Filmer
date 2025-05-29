// Package cache provides storage interface for caching data.
package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	checkConnCtxTimeout = 5 * time.Second // timeout for check conn ctx
	cacheIOCtxTimeout   = 2 * time.Second // timeout for ctx for set/get funcs
)

// Internal interface used only for success connection log output.
type Logger interface {
	Printf(format string, args ...any)
}

// Storage is like fiber.Storage.
type Storage interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
}

// Storage implementation.
type redisStorage struct {
	client *redis.Client
}

func NewStorage(connString string, logger Logger) (Storage, error) {
	// connect to redis
	client, err := createRedisClient(connString)
	if err != nil {
		return nil, err
	}

	logger.Printf("Process %d successfully connected to redis", os.Getpid())
	return &redisStorage{
		client: client,
	}, nil
}

// Get gets the value for the given key.
func (s *redisStorage) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
	defer cancel()
	value, err := s.client.Get(ctx, key).Bytes()

	// if NOT "Not found" error
	if err != nil && !redis.HasErrorPrefix(err, "redis: nil") {
		return nil, fmt.Errorf("get value: %w", err)
	}
	return value, nil
}

// Set stores the given value for the given key along
// with an expiration value, 0 means no expiration.
func (s *redisStorage) Set(key string, val []byte, exp time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
	defer cancel()
	if err := s.client.Set(ctx, key, val, exp).Err(); err != nil {
		return fmt.Errorf("set value: %w", err)
	}
	return nil
}

// Delete deletes the value for the given key.
func (s *redisStorage) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
	defer cancel()
	if err := s.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("delete value: %w", err)
	}
	return nil
}

// Reset deletes all keys.
func (s *redisStorage) Reset() error {
	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
	defer cancel()
	if err := s.client.FlushAll(ctx).Err(); err != nil {
		return fmt.Errorf("reset storage: %w", err)
	}
	return nil
}

// Close closes the storage.
func (s *redisStorage) Close() error {
	if err := s.client.Close(); err != nil {
		return fmt.Errorf("close client: %w", err)
	}
	return nil
}

// createRedisClient opens redis connection by given connection string.
func createRedisClient(connString string) (*redis.Client, error) {
	// create new client
	client := redis.NewClient(&redis.Options{
		Addr: connString,
		DB:   0,
	})
	// redis requests context
	ctx, cancel := context.WithTimeout(context.Background(), checkConnCtxTimeout)
	defer cancel()

	// check connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("connect to redis: %w", err)
	}
	return client, nil
}
