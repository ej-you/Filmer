package redis

import (
	"context"
	"fmt"
	"time"

	goRedis "github.com/redis/go-redis/v9"
	fiber "github.com/gofiber/fiber/v2"

	"server/settings"
)


// помещение токена в список blacklisted в кэше
func SetBlacklistedToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	key := "token:blacklisted:" + token
	// установка ключа-значения в redis со временем просрочки как у токена
	err := GetRedisClient().Set(ctx, key, "true", settings.TokenExpiredTime).Err()
	if err != nil {
	    return fmt.Errorf("token %q: %w", token, fiber.NewError(500, "failed to add token to blacklist: " + err.Error()))
	}
	return nil
}

// true, если токен в списке blacklisted в кэше
func GetBlacklistedToken(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	key := "token:blacklisted:" + token
	// получение значения из redis по ключу
	isBlacklisted, err := GetRedisClient().Get(ctx, key).Bool()
	if err != nil {
		// такого ключа нет в кэше - токен не в чёрном списке
		if goRedis.HasErrorPrefix(err, "redis: nil") {
			return false, nil
		}
	    return false, fmt.Errorf("token %q: find in blacklist: %w", token, fiber.NewError(500, "failed to get token: " + err.Error()))
	}
	return isBlacklisted, nil
}
