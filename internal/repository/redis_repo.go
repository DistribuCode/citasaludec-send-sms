package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func StoreSMS(redisHost, to, message string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
	defer rdb.Close()

	ctx := context.Background()
	key := fmt.Sprintf("sms:%s", to)
	return rdb.Set(ctx, key, message, 0).Err()
}
