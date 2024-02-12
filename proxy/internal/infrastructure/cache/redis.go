package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedisClient(host, port string) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping().Result()

	return client, err
}
