package redis

import (
	"context"
	"os"
	"sync"
	"time"

	goRedis "github.com/redis/go-redis/v9"

	"server/settings"
)


var redisClient *goRedis.Client
var once sync.Once


// получение клиента redis
func GetRedisClient() *goRedis.Client {
	once.Do(func() {
		settings.InfoLog.Printf("Process %d is connecting to redis on %s...", os.Getpid(), settings.RedisAddr)

		// создание нового клиента
		redisClient = goRedis.NewClient(&goRedis.Options{
			Addr: settings.RedisAddr,
			DB: 0,
		})
		// контекст для выполнения запросов к redis
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// проверка подключения
		pong, err := redisClient.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		settings.InfoLog.Printf("Process %d successfully connected to redis: PING - %s", os.Getpid(), pong)		
	})
	return redisClient
}
