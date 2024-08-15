package cache

import (
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	"user_system/conf"
)

type RedisCache struct {
}

func Setup() {
	conf.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisConf.Host, conf.RedisConf.Port),
		Password: conf.RedisConf.Password, // no password set
		DB:       conf.RedisConf.Db,       // use default DB
	})

	//连接redis
	_, err := conf.RedisClient.Ping(conf.CtxRedis).Result()
	//判断连接是否成功
	if err != nil {
		log.Panicf("redis连接失败:%v", err)
	}
}

// Set 设置
func (r RedisCache) Set(key string, value interface{}, expiration int) {
	v, err := json.Marshal(value) //value是一个空接口类型,里面可以是字符串,切片,结构体,所以转成json保存
	if err == nil {
		//RedisDb:调用redisCore.go中的RedisDb
		conf.RedisClient.Set(conf.CtxRedis, key, string(v), time.Second*time.Duration(expiration))
	}
}

// Get 获取
func (r RedisCache) Get(key string, obj interface{}) bool {
	valueStr, err1 := conf.RedisClient.Get(conf.CtxRedis, key).Result()
	if err1 == nil && valueStr != "" {
		err2 := json.Unmarshal([]byte(valueStr), obj)
		return err2 == nil
	}
	return false
}

func (r RedisCache) Del(key string) bool {
	_, err := conf.RedisClient.Del(conf.CtxRedis, key).Result()
	return err == nil
}
