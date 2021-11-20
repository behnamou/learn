package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// err := client.Set(ctx, "key1", "val1", 0).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err := client.Set(ctx, "key2", "val2", 0).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	client.Del(ctx, "key2")

	ans, _ := client.Keys(ctx, "key*").Result()
	fmt.Println("key", ans)

}
