package config

import "fmt"

type CommonConf struct {
}

type MysqlConf struct {
}

type RpcConf struct {
}

type RedisConf struct {
}

var (
	conf CommonConf
	mysqlConf MysqlConf
	rpcConf RpcConf
	redisConf RedisConf
)

func GetConf() CommonConf {
	return conf
}

func GetMysqlConf() MysqlConf {
	return mysqlConf
}

func GetRpcConf() RpcConf {
	return rpcConf
}

func GetRedisConf() RedisConf {
	return redisConf
}

func init() {
	// config
	fmt.Println("config init!!!")
	//if _, err := toml.DecodeFile("test.toml", &Conf); err != nil {
	//	panic(err)
	//}
}
