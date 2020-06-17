package main

import (
	"fmt"
	_ "gin-demo/config"
	"gin-demo/route"
	_ "gin-demo/utils/validater"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init conf!!!")
	gin.SetMode("debug")
}

func main() {
	engine := gin.New()

	// 路由设置
	route.SetupRouter(engine)
	engine.Run(":6666")
}
