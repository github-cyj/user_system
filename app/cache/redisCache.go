package cache

import (
	"encoding/json"
	"time"
	"user_system/config"
)

type RedisCache struct {
}

// Set 设置
func (r RedisCache) Set(key string, value interface{}, expiration int) {
	v, err := json.Marshal(value) //value是一个空接口类型,里面可以是字符串,切片,结构体,所以转成json保存
	if err == nil {
		//RedisDb:调用redisCore.go中的RedisDb
		config.RedisDb.Set(config.CtxRedis, key, string(v), time.Second*time.Duration(expiration))
	}
}

// Get 获取
func (r RedisCache) Get(key string, obj interface{}) bool {
	valueStr, err1 := config.RedisDb.Get(config.CtxRedis, key).Result()
	if err1 == nil && valueStr != "" {
		err2 := json.Unmarshal([]byte(valueStr), obj)
		return err2 == nil
	}
	return false
}

func (r RedisCache) Del(key string) bool {
	_, err := config.RedisDb.Del(config.CtxRedis, key).Result()
	return err == nil
}
