package credis

import "github.com/go-redis/redis"

func RedisInit(addr string) (*redis.Client, error) {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "XXXX",
		DB:       0,
		PoolSize: 10,
		MaxRetries: 5,
		MinIdleConns: 5,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		return nil, err
	} else {
		return redisdb, nil
	}
}