package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func Add(c *gin.Context) {
	args := c.Keys["args"].(*AddEntity)

	//i := 0
	//t := 1/i
	//fmt.Println(t)
	res := args
	fmt.Println(res)
	// 业务处理...

	c.JSON(200, gin.H{"code": 200, "msg": "", "data": gin.H{"id": res.Id, "name": res.Name}})
}

func Get(c *gin.Context) {

}
