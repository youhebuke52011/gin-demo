package user

import (
	"gin-demo/client"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Add(c *gin.Context) {
	var (
		args   = c.Keys["args"].(*AddEntity)
		result = gin.H{"code": 200, "msg": "ok", "data": gin.H{}}
	)

	if err := client.GetMysqlCli().Demo.Table("user").Create(args).Error; err != nil {
		log.WithFields(log.Fields{"err": err}).Error("gorm")
		result["code"] = 500
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	var (
		args   = c.Keys["args"].(*GetEntity)
		result = gin.H{"code": 200, "msg": "ok", "data": gin.H{}}
		user = User{}
	)

	if err := client.GetMysqlCli().Demo.Table("user").Where("id = ?", args.Id).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.WithFields(log.Fields{"err": err}).Error("gorm")
		}
		c.JSON(http.StatusOK, result)
		return
	}
	result["data"] = gin.H{"id": user.Id, "name": user.Name}
	c.JSON(http.StatusOK, result)
}
