package core

import (
	"gin-demo/core/middleware/exception"
	"gin-demo/core/middleware/logger"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	// middleware
	Engine.Use(exception.SetUp(), logger.SetUp())
}

