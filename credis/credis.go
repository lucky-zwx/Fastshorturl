package credis

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/go-redis/redis"
)

var Cfg, err = goconfig.LoadConfigFile("conf/runconf.ini")

func RedisInit() *redis.Client {

	if err != nil {
		fmt.Println("未找到配置文件")
		return nil
	}

	Addr, err := Cfg.GetValue("redis", "Addr")
	Password, err := Cfg.GetValue("redis", "Password")
	DB, err := Cfg.Int("redis", "DB")
	PoolSize, err := Cfg.Int("redis", "PoolSize")
	MaxRetries, err := Cfg.Int("redis", "MaxRetries")
	MinIdleConns, err := Cfg.Int("redis", "MinIdleConns")

	redisdb := redis.NewClient(&redis.Options{
		Addr:         Addr,
		Password:     Password,
		DB:           DB,
		PoolSize:     PoolSize,
		MaxRetries:   MaxRetries,
		MinIdleConns: MinIdleConns,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return nil
	} else {
		return redisdb
	}
}
