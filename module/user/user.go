package user

import (
	"fmt"
	"gin-demo/utils/validater"
	"reflect"

	//"gin-demo/utils/validater"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func bind(s interface{}, c *gin.Context) (interface{}, error) {
	//b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindBodyWith(s, binding.JSON);err != nil {
	//if err := c.ShouldBindWith(s, b); err != nil {
		return nil, err
	}
	return s, nil
}

type ProductAdd struct {
	Id   int    `json:"id" form:"id" binding:"required,gt=18"`
	Name string `json:"name" form:"name" binding:"required"`
	//Name string `form:"name" json:"name" validate:"required,enum"`
}

func Add(c *gin.Context) {
	typs := c.Keys["type"].(reflect.Type)
	//typs := reflect.TypeOf(ProductAdd{})
	argv := reflect.New(typs)
	args := argv.Interface()
	err := c.ShouldBindBodyWith(args, binding.JSON)
	if err == nil {
		c.Set("args", args)
	}

	//fmt.Println(c.Request.Header)

	res := args.(*ProductAdd)
	fmt.Println(res)
	// 业务处理...

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"id": res.Id, "name": res.Name}})
}

func Get(c *gin.Context) {

}

func init() {
	v := binding.Validator.Engine().(*validator.Validate)

	err := v.RegisterValidation("enum", validater.ValidEnum)
	fmt.Println(err)
}