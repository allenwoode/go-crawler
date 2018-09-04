package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

func NewRedisClient(addr string, passwd string, db int) *redis.Client  {
	return redis.NewClient(&redis.Options{
		Addr: addr,
		Password: passwd,
		DB: db,
	})
}

func main()  {
	c := NewRedisClient("localhost:6379", "", 0)

	pong, err := c.Ping().Result()
	fmt.Println(pong, err)

	key := "ZH"
	err = c.Set(key, "中国", 5 * time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, _ := c.Get(key).Result()
	fmt.Println(key, val)
}
