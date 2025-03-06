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


// интерфейс Cache для кэша приложения
type Cache interface {
	Set(key string, value any, expiration time.Duration) error
	GetBool(key string) (bool, error)
}


// реализация кэша через redis
type redisCache struct {
	cfg		*config.Config
	redis	*goRedis.Client
}

var redisCacheInstance redisCache
var once sync.Once

// конструктор для типа интерфейса Cache
func NewCache(cfg *config.Config, log logger.Logger) Cache {
	once.Do(func() {
		log.Infof("Process %d is connecting to redis on %s...", os.Getpid(), cfg.Cache.ConnString)

		// создание нового клиента
		redisCacheInstance.redis = goRedis.NewClient(&goRedis.Options{
			Addr: cfg.Cache.ConnString,
			DB: 0,
		})
		// контекст для выполнения запросов к redis
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// проверка подключения
		pong, err := redisCacheInstance.redis.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		log.Infof("Process %d successfully connected to redis: PING - %s", os.Getpid(), pong)		
		
		redisCacheInstance.cfg = cfg
	})
	return &redisCacheInstance
}

// установка ключа-значения в redis с переданным временем просрочки
func (this redisCache) Set(key string, value any, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return this.redis.Set(ctx, key, value, expiration).Err()
}

// получение bool значения из redis по ключу
func (this redisCache) GetBool(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	value, err := this.redis.Get(ctx, key).Bool()

	// если НЕ ошибка ненахождения значения по ключу
	if err != nil && !goRedis.HasErrorPrefix(err, "redis: nil") {
		return false, err
	}
	return value, nil
}
