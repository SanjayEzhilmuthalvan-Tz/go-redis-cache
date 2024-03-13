package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Author{Name: "sanjay", Age: 20})
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set("id4", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id4").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
