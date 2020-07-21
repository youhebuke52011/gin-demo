package core

import (
	"gin-demo/core/middleware"
	"gin-demo/module/singer"
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

	engine.GET("/v2/test/gz", gzip.Gzip(gzip.BestSpeed), tg.TGzip)
	engine.GET("/v2/singer/list", singer.List)
	engine.GET("/v2/singer/detail", singer.Detail)
	engine.POST("/v2/singer/add", HandleCore(reflect.TypeOf(singer.AddEntity{}), singer.Add, []CheckHandle{middleware.BindParam}))
	engine.POST("/v2/singer/update", HandleCore(reflect.TypeOf(singer.AddEntity{}), singer.Update, []CheckHandle{middleware.BindParam}))
	engine.GET("/v2/singer/delete", singer.Delete)
	engine.POST("/v2/user/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "login",
			"code": 200,
		})
	})
	engine.GET("/v2/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": gin.H{
				"id": 1,
				"name": "a",
				"avatar": "a",
				"introduction": "a",
				"roles": []string{"admin"},
			},
			"code": 200,
		})
	})
}
