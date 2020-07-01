package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

func BindParam(c *gin.Context) bool {
	var (
		typ = c.Keys["type"].(reflect.Type)
		etip string
		tips = []string{}
		argv = reflect.New(typ)
		args = argv.Interface()
		err error
	)
	b := binding.Default(c.Request.Method, c.ContentType())
	err = c.ShouldBindWith(args, b)
	if err == nil {
		c.Set("args", args)
		return true
	}
	switch err.(type) {
	case validator.ValidationErrors:
		fmt.Println(1)
		errs := err.(validator.ValidationErrors)
		for _, f := range errs {
			tip := fmt.Sprintf("field:%s,rule:%s", strings.ToLower(f.Field()), f.Tag())
			tips = append(tips, tip)
		}
	case *json.UnmarshalTypeError:
		fmt.Println(2)
		errs := err.(*json.UnmarshalTypeError)
		tip := fmt.Sprintf("field:%s,rule:type", strings.ToLower(errs.Field))
		tips = append(tips, tip)
	}
	if len(tips) == 0 {
		etip = "invalid json format"
	} else {
		etip = strings.Join(tips, "|")
	}
	//fmt.Println(etip)
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": etip, "data": gin.H{}})
	return false
}
