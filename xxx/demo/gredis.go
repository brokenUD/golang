package demotest

import (
	"context"
	"fmt"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func stringg(ctx context.Context, client *redis.Client){
	key := "aaa"
	value := "2131"
	err := client.Set(ctx, key, value, 1*time.Second).Err()
	checkErr(err)

	time.Sleep(2*time.Second)
	v2, err := client.Get(ctx, key).Result()
	checkErr(err)
	fmt.Println(v2)

	client.Del(ctx, key)
}

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RedisTest(){
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	ctx := context.TODO()
	stringg(ctx, client)
}
