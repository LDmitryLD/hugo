package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const defaultReconnectionTimeout = 5

func NewRedisClient(host, port string) (*redis.Client, error) {

	// client := redis.NewClient(&redis.Options{
	// 	Addr: "redis:6379",
	// })

	// addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err == nil {
		return client, nil
	}

	log.Println("error with starting Redis server:", err)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(defaultReconnectionTimeout))

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("cache connection failed after %d timeout: %s", defaultReconnectionTimeout, err.Error())

		case <-ticker.C:
			err := client.Ping(ctx).Err()
			if err == nil {
				return client, nil
			}
			log.Println("error when starting Redis server")
		}
	}
}
