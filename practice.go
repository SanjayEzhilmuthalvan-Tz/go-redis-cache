package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	type Person struct {
		Name string
		Age  int
	}

	err := client.HSet(ctx, "STRUCT", "name", "sj", "age", "20").Err()
	if err != nil {
		panic(err)
	}

	val1, err1 := client.HGetAll(ctx, "STRUCT").Result()
	if err1 != nil {
		panic(err)
	}

	fmt.Println("struct data", val1)
}
