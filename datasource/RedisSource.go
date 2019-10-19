package datasource

import (
	"StreamChannelSwitch/config"
	"github.com/go-redis/redis"
)

type RedisSouce struct {
	client     *redis.Client
	config     *config.Config
	keytimepit int8
}

func NewRedisSouce(config *config.Config) *RedisSouce {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	return &RedisSouce{config: config, client: client}
}

func (r *RedisSouce) AddClass(sid string, value interface{}) error {
	key, err := r.client.Get(sid).Result()
	if err != nil {
		return err
	}

	if len(key) != 0 {

	}

	return nil
}

func (r *RedisSouce) DelClass(sid string, value interface{}) error {

	return nil
}

func (r *RedisSouce) GetClass(sid string) (error, value interface{}) {

	return nil, nil
}
