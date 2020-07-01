package core

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type CheckHandle func(ctx *gin.Context) bool
type FunHandle func(ctx *gin.Context)

func HandleCore(typ reflect.Type, handle FunHandle, checks []CheckHandle) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Set("type", typ)
		for _, cb := range checks {
			if !cb(c) {
				return
			}
		}
		handle(c)
	}
}




