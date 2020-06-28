package main

import (
	"fmt"
	_ "gin-demo/config"
	"gin-demo/core"
	"gin-demo/core/middleware"
	_ "gin-demo/common/validater"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

var Engine *gin.Engine

func init() {
	fmt.Println("init main!!!")
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.SetOutput(os.Stdout)
}

func main() {
	Engine = gin.New()

	// 路由设置
	Engine.Use(middleware.Exception(), middleware.Logger())
	core.SetupRouter(Engine)
	Engine.Run(":6666")
}
