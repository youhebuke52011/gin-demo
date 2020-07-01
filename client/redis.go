package client

import (
	"fmt"
	"gin-demo/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
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
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s/%d", addr, db))
	if err != nil {
		log.WithFields(log.Fields{"database": db, "error": err}).Error("redis connect error")
		panic(err)
	}
	rdb := redis.NewClient(opt)
	log.WithFields(log.Fields{"database": db}).Info("redis connect success")
	return rdb
}

func init() {
	conf := config.GetConf()
	demoConf := conf.GetStringMapString("redis.demo")
	redisCli = &RedisCli{
		Demo: getRdbConn(demoConf["database"], "", demoConf["password"], demoConf["addr"]),
	}
}
