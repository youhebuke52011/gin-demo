package main

import (
	"fmt"
	"gin-demo/route"
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
	engine.Run(":9999")
}

