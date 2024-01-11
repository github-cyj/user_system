package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"user_system/utils"
)

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
}

// 全局使用,就需要把定义成公有的
var (
	CtxRedis = context.Background()
	RedisDb  *redis.Client
)
var RedisConfig = &Redis{}

func (r Redis) Setup() {
	utils.LoadConfig("redis", RedisConfig)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port),
		Password: RedisConfig.Password, // no password set
		DB:       RedisConfig.Db,       // use default DB
	})

	//连接redis
	_, err := RedisDb.Ping(CtxRedis).Result()
	//判断连接是否成功
	if err != nil {
		log.Panicf("redis连接失败:%v", err)
	}

}
