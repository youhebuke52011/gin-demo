package middleware

import (
	"github.com/gin-gonic/gin"
)


func TestMiddleWare() gin.HandlerFunc {
	//log.WithFields(log.Fields{"test": "tt111"}).Info("test middleware!")
	return func(c *gin.Context) {
		//log.WithFields(log.Fields{"test": "tt222"}).Info("test middleware!")
		c.Next()
		//log.WithFields(log.Fields{"test": "tt333"}).Info("test middleware!")
	}
}
