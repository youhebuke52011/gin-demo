package main

import (
	_ "gin-demo/common/validater"
	_ "gin-demo/configs"
	"gin-demo/core"
	"gin-demo/core/middleware"
	"gin-demo/module/cron"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var Engine *gin.Engine

func init() {
	log.SetLevel(log.InfoLevel)
	//设置日志格式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	go func() {
		// net/http/pprof 注册是的默认的mux
		http.ListenAndServe(":6060", nil)
	}()
}

func cancel() {
	//client.GetMysqlCli().Close()
	//client.GetRedisCli().Close()
	cron.Close()
}

func main() {
	Engine = gin.New()
	defer cancel()

	// 路由设置
	//  middleware.Sign()
	Engine.Use(middleware.Exception(), middleware.Logger())
	core.SetupRouter(Engine)
	Engine.Run(":6666")

}
