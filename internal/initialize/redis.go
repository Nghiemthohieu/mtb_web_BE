package initialize

import (
	"context"
	"fmt"
	"mtb_web/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis Connect Error: ", zap.Error(err))
		panic(err)
	}
	fmt.Println("InitRedis is runnning")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(context.Background(), "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Set Error: ", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(context.Background(), "score").Result()
	if err != nil {
		fmt.Println("Get Error: ", zap.Error(err))
		return
	}
	global.Logger.Info("value score is: ", zap.String("score", value))
}
