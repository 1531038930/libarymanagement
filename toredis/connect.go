package toredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Rdb *redis.Client
	ctx context.Context
)

func init() {
	ctx = context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	//defer Rdb.Close()
}

func Add(userid uint, token string) error {
	if err := Rdb.Set(ctx, token, userid, 600*time.Second).Err(); err != nil {
		fmt.Println("无法连接到Redis:", err)
		return err
	}
	return nil
}
func Check(token string, id string) error {
	get := Rdb.Get(ctx, token)
	if get.Err() != nil {
		return errors.New("token无效")
	}
	if get.Val() != id {
		return errors.New("id与token不匹配，请重试")
	}
	return nil
}
