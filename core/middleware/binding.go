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
	if c.Request.Method == http.MethodPost {
		err = c.ShouldBindBodyWith(args, binding.JSON)
	} else {
		err = c.ShouldBindWith(args, binding.Form)
	}
	if err == nil {
		c.Set("args", args)
		return true
	}
	switch err.(type) {
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		for _, f := range errs {
			tip := fmt.Sprintf("field:%s,rule:%s", strings.ToLower(f.Field()), f.Tag())
			tips = append(tips, tip)
		}
	case *json.UnmarshalTypeError:
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
	c.JSON(http.StatusBadRequest, gin.H{"code": 400001, "msg": etip, "data": gin.H{}})
	return false
}