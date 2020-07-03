package client

import (
	"fmt"
	"gin-demo/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type MysqlCli struct {
	Demo *gorm.DB
}

var (
	mysqlCli *MysqlCli
)

func GetMysqlCli() *MysqlCli {
	return mysqlCli
}

func (mysql *MysqlCli) Close() {
	mysql.Demo.Close()
}

func getMysqlConn(database, user, password, addr string) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", user, password, addr, database))
	if err != nil {
		log.WithFields(log.Fields{"database": database, "error": err}).Error("mysql connect error")
		panic(err)
	}
	log.WithFields(log.Fields{"database": database}).Info("mysql connect success")

	return db
}

func init() {
	conf := configs.GetConf()
	demoConf := conf.GetStringMapString("mysql.demo")
	mysqlCli = &MysqlCli{
		Demo: getMysqlConn(demoConf["database"], demoConf["user"], demoConf["password"], demoConf["addr"]),
	}
}
