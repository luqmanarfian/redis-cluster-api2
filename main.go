package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// ip yang akan digunakan  "34.68.30.186:7000", "35.202.198.198:7000", "34.172.44.137:7000", "34.67.142.60:7000", "34.66.14.86:7000", "35.222.248.41:7000"
func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"34.68.30.186:7000", "35.202.198.198:7000", "34.172.44.137:7000", "34.67.142.60:7000", "35.222.248.41:7000"},

		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})
	ctx := context.Background()

	err := rdb.ForEachMaster(ctx, func(ctx context.Context, master *redis.Client) error {
		return master.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	val, err := rdb.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val.(string))

}
