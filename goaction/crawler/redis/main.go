package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func pong(client *redis.Client)  {
	pong, err := client.Ping().Result()

	fmt.Println(pong, err)
}

func work(client *redis.Client)  {
	err := client.Set("china", "中国", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("china").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("china:", val)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong(client)
	work(client)

	val, err := client.Get("american").Result()
	if err == redis.Nil {
		fmt.Println("key 'american' does not exist")
	} else if err == nil {
		panic(err)
	} else {
		fmt.Println("american:", val)
	}
}
