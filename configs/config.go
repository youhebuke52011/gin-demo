package configs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//type CommonConf struct {
//}
//
//type MysqlConf struct {
//}
//
//type RpcConf struct {
//}
//
//type RedisConf struct {
//}
//
//var (
//	conf CommonConf
//	mysqlConf MysqlConf
//	rpcConf RpcConf
//	redisConf RedisConf
//)
//
//func GetConf() CommonConf {
//	return conf
//}
//
//func GetMysqlConf() MysqlConf {
//	return mysqlConf
//}
//
//func GetRpcConf() RpcConf {
//	return rpcConf
//}
//
//func GetRedisConf() RedisConf {
//	return redisConf
//}

var (
	conf *viper.Viper
	envConf *viper.Viper
)

func GetConf() *viper.Viper {
	return conf
}

func GetEnvConf() *viper.Viper {
	return envConf
}

func readConfig(configName string) *viper.Viper {
	v := viper.New()

	v.SetConfigName(configName) // name of config file (without extension)
	v.AddConfigPath("./configs/")        // optionally look for config in the working directory
	err := v.ReadInConfig()     // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

func init() {
	// config
	conf = readConfig("conf")
	log.WithFields(log.Fields{}).Info("config is ready")
	
}
