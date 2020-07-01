package main

import (
	"gin-demo/client"
	_ "gin-demo/common/validater"
	_ "gin-demo/config"
	"gin-demo/core"
	"gin-demo/core/middleware"
	"gin-demo/module/cron"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

var Engine *gin.Engine

func init() {
	log.SetLevel(log.DebugLevel)
	//log.SetOutput(os.Stdout)
	go func() {
		// net/http/pprof 注册是的默认的mux
		http.ListenAndServe(":6060", nil)
	}()
}

func cancel() {
	client.GetMysqlCli().Close()
	client.GetRedisCli().Close()
	cron.Close()
}

func main() {
	Engine = gin.New()
	defer cancel()

	// 路由设置
	Engine.Use(middleware.Exception(), middleware.Logger())
	core.SetupRouter(Engine)
	Engine.Run(":6666")

}
