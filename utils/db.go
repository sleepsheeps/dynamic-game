package utils

import (
	"errors"
	"github.com/go-redis/redis"
)

var (
	cache *Cache
)

type Cache struct {
	*redis.Client
}

func (c *Cache) Init(addr string) error {
	c.Client = redis.NewClient(&redis.Options{Addr: addr})
	return nil
}

func (c *Cache) GetString(key string) (string, error) {
	cmd := redis.NewStringCmd("get", key)
	c.Process(cmd)
	return cmd.Result()
}

func (c *Cache) SetString(key string, val string) error {
	cmd := redis.NewStringCmd("set", key, val)
	return c.Process(cmd)
}

func (c *Cache) Del(key string) error {
	cmd := redis.NewStringCmd("del", key)
	return c.Process(cmd)
}

func InitCache(addr string) error {
	if cache == nil {
		cache = new(Cache)
	}
	return cache.Init(addr)
}

func GetString(key string) (string, error) {
	if cache == nil {
		return "", errors.New("cache need init")
	}
	return cache.GetString(key)
}

func SetString(key, val string) error {
	if cache == nil {
		return errors.New("cache need init")
	}
	return cache.SetString(key, val)
}

func Del(key string) error {
	if cache == nil {
		return errors.New("cache need init")
	}
	return cache.Del(key)
}
