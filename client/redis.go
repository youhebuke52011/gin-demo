package client

import (
	"gin-demo/config"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisCli struct {
	Demo *redis.Client
}

var (
	redisCli *RedisCli
)

func GetRedisCli() *RedisCli {
	return redisCli
}

func (r *RedisCli) Close() {
	r.Demo.Close()
}

func getRdbConn(database, user, password, addr string) *redis.Client {
	db, _ := strconv.Atoi(database)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return rdb
}

func init() {
	conf := config.GetConf()
	demoConf := conf.GetStringMapString("redis.demo")
	redisCli = &RedisCli{
		Demo: getRdbConn(demoConf["database"], "", demoConf["password"], demoConf["addr"]),
	}
}
