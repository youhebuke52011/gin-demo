package main

import (
	"fmt"
	"gin-demo/route"
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
	route.SetupRouter(engine)
	engine.Run(":6666")
}

