package core

import (
	"gin-demo/core/middleware"
	tg "gin-demo/module/test"
	"gin-demo/module/user"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"reflect"
)

func SetupRouter(engine *gin.Engine) {

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

	userGroup := engine.Group("/user")
	{
		userGroup.GET("", HandleCore(
			reflect.TypeOf(user.GetEntity{}), user.Get, []CheckHandle{middleware.BindParam, middleware.Sign, middleware.RateLimit}))

		userGroup.POST("", middleware.TestMiddleWare(), HandleCore(
			reflect.TypeOf(user.AddEntity{}), user.Add, []CheckHandle{middleware.BindParam, middleware.Sign}))
	}

	engine.GET("/test/gz", gzip.Gzip(gzip.DefaultCompression), tg.TGzip)
}
