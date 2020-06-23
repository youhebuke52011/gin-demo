package main

import (
	"fmt"
	_ "gin-demo/config"
	"gin-demo/core"
	_ "gin-demo/utils/validater"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("init conf!!!")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	gin.SetMode("debug")
}

func main() {
	engine := gin.New()

	// 路由设置
	core.SetupRouter(engine)
	engine.Run(":6666")
}
