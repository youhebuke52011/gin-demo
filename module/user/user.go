package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func bind(s interface{}, c *gin.Context) (interface{}, error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(s, b); err != nil {
		return nil, err
	}
	return s, nil
}

type ProductAdd struct {
	Id   int    `json:"id" validate:"required,enum=1"`
	Name string `json:"name" validate:"required"`
	//Name string `form:"name" json:"name" validate:"required,enum"`
}

func Add(c *gin.Context) {
	args, err := bind(&ProductAdd{}, c)
	if err != nil {
		c.JSON(400, nil)
		fmt.Println(err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&args); err != nil {
		fmt.Printf("validate add:%v\n", err)
		c.JSON(200, gin.H{"code": 400, "data": gin.H{}})
		return
	}

	res := args.(*ProductAdd)
	fmt.Println(res)
	// 业务处理...

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"id": res.Id, "name": res.Name}})
}

func Get(c *gin.Context) {

}
