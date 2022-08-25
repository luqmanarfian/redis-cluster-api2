package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// ip yang akan digunakan  "34.68.30.186:7000", "35.202.198.198:7000", "34.172.44.137:7000", "34.67.142.60:7000", "34.66.14.86:7000", "35.222.248.41:7000"
func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"34.68.30.186:7000", "35.202.198.198:7000", "34.172.44.137:7000", "34.67.142.60:7000", "34.66.14.86:7000", "35.222.248.41:7000"},

		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})
	ctx := context.Background()

	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}

}
