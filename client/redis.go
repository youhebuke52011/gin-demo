package client

import (
	"github.com/go-redis/redis/v8"
)

type RedisCli struct {
	demo *redis.Client
}

var (
	redisCli *RedisCli
)

func GetRedisCli() *RedisCli {
	return redisCli
}

func (r *RedisCli) Close() {
	r.demo.Close()
}

func getRdbConn(database, user, password, addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func init() {
	redisCli = &RedisCli{
		demo: getRdbConn("", "", "", ""),
	}
}
