package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"reflect"
)

func BindParam(c *gin.Context) bool {
	typs := c.Keys["type"].(reflect.Type)
	//typs := reflect.TypeOf(ProductAdd{})
	argv := reflect.New(typs)
	args := argv.Interface()
	err := c.ShouldBindBodyWith(args, binding.JSON)
	if err == nil {
		c.Set("args", args)
	}
	return false
}