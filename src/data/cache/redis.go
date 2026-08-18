package cache

import (
	"fmt"
	"time"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config)error{

	var rdsConfig = &cfg.Redis

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",rdsConfig.Host,rdsConfig.Port),
		Password: rdsConfig.Password,
		DB: 0,
		DialTimeout: rdsConfig.DialTimeout * time.Second,
		ReadTimeout: rdsConfig.ReadTimeout * time.Second,
		WriteTimeout: rdsConfig.WriteTimeout * time.Second,
		PoolSize: rdsConfig.PoolSize,
		PoolTimeout: rdsConfig.PoolTimeout,
		IdleTimeout: 500 * time.Millisecond,
		IdleCheckFrequency: rdsConfig.IdleCheckFrequency * time.Millisecond,
	})
	_,err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client{
	return redisClient
}

func CloseRedis(){
	redisClient.Close()
}