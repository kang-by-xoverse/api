package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() (*redis.Client, func()) {
	rdb := redis.NewClient(&redis.Options{
		Username: GetEnv("REDIS_USERNAME", "default"),
		Addr:     GetEnv("REDIS_ADDRESS", "localhost"),
		Password: GetEnv("REDIS_PASSWORD", "secret"),
		DB:       0,
	})

	return rdb, func() {
		rdb.Close()
	}
}

func Emit(rdb *redis.Client, event string, data interface{}) {
	var ctx = context.Background()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// rdb.Set(ctx, event, string(jsonData), time.Hour)
	rdb.Publish(ctx, event, string(jsonData))
}