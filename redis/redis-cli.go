package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/my0sot1s/social/utils"
)

type RedisCli struct {
	client   *redis.Client
	duration time.Duration
}

func (rc *RedisCli) Config(redisHost, redisDb, redisPass string) error {

	rc.client = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})
	rc.duration = 60 * 10 * time.Second // 10 phut

	pong, err := rc.client.Ping().Result()
	fmt.Println(pong, err)
	utils.Log("ಠ‿ಠ Redis connected ಠ‿ಠ")
	return nil
}
func (rc *RedisCli) Close() error {
	return rc.Close()
}

func (rc *RedisCli) SetValue(key string, value string, expiration time.Duration) error {
	return rc.client.Set(key, value, expiration*time.Second).Err()
}

func (rc *RedisCli) GetValue(key string) (string, error) {
	val, err := rc.client.Get(key).Result()
	return val, err
}
