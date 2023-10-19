package toredis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Con(userid uint, token string) error {
	ctx := context.Background()
	cli := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	defer cli.Close()
	if err := cli.Set(ctx, token, userid, 600*time.Second).Err(); err != nil {
		fmt.Println("无法连接到Redis:", err)
		return err
	}
	return nil
}
