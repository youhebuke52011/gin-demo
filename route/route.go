package route

import (
	"gin-demo/route/middleware/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	// middleware
	engine.Use(logger.SetUp())

	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "请求方法不存在",
		})
	})

	// ping
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
}
