package client

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

type MysqlCli struct {
	demo *gorm.DB
}

var (
	mysqlCli *MysqlCli
)

func GetMysqlCli() *MysqlCli {
	return mysqlCli
}

func (mysql *MysqlCli) Close() {
	mysql.demo.Close()
}

func getMysqlConn(database, user, password, addr string) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", user, password, addr, database))
	if err != nil {
		log.WithFields(log.Fields{"database": database, "error": err}).Error("mysql connnect error")
		return nil
	}
	log.WithFields(log.Fields{"database": database}).Info("mysql connnect success")

	return db
}

func init() {
	mysqlCli = &MysqlCli{
		demo: getMysqlConn("", "", "", ""),
	}
}
