package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
)

type RedisCli struct {
	*redis.Client
}

type Config struct {
	Host     string
	Port     int
	DB       int
	Password string
}

func New(c *Config) *RedisCli {
	var r = new(RedisCli)
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		DB:       c.DB, // use default DB
		Password: c.Password,
	})
	return r
}

func NewByS(addr, password string) *RedisCli {
	var r = new(RedisCli)
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0, // use default DB
		PoolSize: 100,
		Password: password,
	})
	return r
}

// HSetFromStruct 把struct按hash结构存入redis
func (r *RedisCli) HSetFromStruct(ctx context.Context, key string, data interface{}) *redis.IntCmd {
	mapData := make(map[string]string)
	d := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < d.NumField(); i++ {
		mapData[d.Field(i).Name] = fmt.Sprint(v.Field(i).Interface())
	}
	return r.HSet(ctx, key, mapData)
}

// HSetFromStructByPip  使用pipeline把struct按hash结构存入redis
func (r *RedisCli) HSetFromStructByPip(ctx context.Context, pip *redis.Pipeliner, key string, data interface{}) *redis.IntCmd {
	mapData := make(map[string]string)
	d := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < d.NumField(); i++ {
		mapData[d.Field(i).Name] = fmt.Sprint(v.Field(i).Interface())
	}
	return (*pip).HSet(ctx, key, mapData)
}
