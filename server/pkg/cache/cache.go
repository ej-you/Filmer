package cache

import (
	"context"
	"os"
	"sync"
	"time"

	goRedis "github.com/redis/go-redis/v9"

	"Filmer/server/pkg/logger"
	"Filmer/server/config"
)


// Cache interface for app cache
type Cache interface {
	Set(key string, value any, expiration time.Duration) error
	GetBool(key string) (bool, error)
}


// Cache implementation through Redis
type redisCache struct {
	cfg		*config.Config
	redis	*goRedis.Client
}

var redisCacheInstance redisCache
var once sync.Once

// Cache constructor
func NewCache(cfg *config.Config, log logger.Logger) Cache {
	once.Do(func() {
		log.Infof("Process %d is connecting to redis on %s...", os.Getpid(), cfg.Cache.ConnString)

		// create new client
		redisCacheInstance.redis = goRedis.NewClient(&goRedis.Options{
			Addr: cfg.Cache.ConnString,
			DB: 0,
		})
		// redis requests context контекст для выполнения запросов к redis
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// check connection
		pong, err := redisCacheInstance.redis.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		log.Infof("Process %d successfully connected to redis: PING - %s", os.Getpid(), pong)		
		
		redisCacheInstance.cfg = cfg
	})
	return &redisCacheInstance
}

// set key-value into redis with expiration time
func (this redisCache) Set(key string, value any, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return this.redis.Set(ctx, key, value, expiration).Err()
}

// get bool value from redis with key
func (this redisCache) GetBool(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	value, err := this.redis.Get(ctx, key).Bool()

	// if NOT "Not found" error
	if err != nil && !goRedis.HasErrorPrefix(err, "redis: nil") {
		return false, err
	}
	return value, nil
}
