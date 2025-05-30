package cache

// import (
// 	"context"
// 	"os"
// 	"sync"
// 	"time"

// 	goredis "github.com/redis/go-redis/v9"
// )

// // Cache interface for app cache
// type Cache interface {
// 	Set(key string, value any, expiration time.Duration) error
// 	GetBool(key string) (bool, error)
// 	GetBytes(key string) ([]byte, error)
// }

// // Cache implementation through Redis
// type redisCache struct {
// 	redis *goredis.Client
// }

// var redisCacheInstance redisCache
// var once sync.Once

// // Cache constructor
// func NewCache(connString string, log Logger) Cache {
// 	once.Do(func() {
// 		// create new client
// 		redisCacheInstance.redis = goredis.NewClient(&goredis.Options{
// 			Addr: connString,
// 			DB:   0,
// 		})
// 		// redis requests context контекст для выполнения запросов к redis
// 		ctx, cancel := context.WithTimeout(context.Background(), checkConnCtxTimeout)
// 		defer cancel()

// 		// check connection
// 		pong, err := redisCacheInstance.redis.Ping(ctx).Result()
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Printf("Process %d successfully connected to redis: PING - %s", os.Getpid(), pong)
// 	})
// 	return &redisCacheInstance
// }

// // set key-value into redis with expiration time
// func (rc redisCache) Set(key string, value any, expiration time.Duration) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
// 	defer cancel()
// 	return rc.redis.Set(ctx, key, value, expiration).Err()
// }

// // get bool value from redis with key
// func (rc redisCache) GetBool(key string) (bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
// 	defer cancel()
// 	value, err := rc.redis.Get(ctx, key).Bool()

// 	// if NOT "Not found" error
// 	if err != nil && !goredis.HasErrorPrefix(err, "redis: nil") {
// 		return false, err
// 	}
// 	return value, nil
// }

// // get bytes value from redis with key
// func (rc redisCache) GetBytes(key string) ([]byte, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), cacheIOCtxTimeout)
// 	defer cancel()
// 	value, err := rc.redis.Get(ctx, key).Bytes()

// 	// if NOT "Not found" error
// 	if err != nil && !goredis.HasErrorPrefix(err, "redis: nil") {
// 		return nil, err
// 	}
// 	return value, nil
// }
