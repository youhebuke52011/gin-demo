package middleware

import (
	"crypto/md5"
	"fmt"
	"gin-demo/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func Sign(c *gin.Context) bool {
	var (
		args      = c.Keys["args"]
		vals      = reflect.Indirect(reflect.ValueOf(args))
		appId     = vals.FieldByName("AppID").Int()
		md5cal    = md5.New()
		appMap    = configs.GetConf().GetStringMap("app")
		appSecret = appMap[strconv.Itoa(int(appId))].(map[string]interface{})["appsecret"]
	)
	if c.Request.Method == http.MethodPost {
		md5cal.Write(c.Keys[gin.BodyBytesKey].([]byte))
	} else {
		md5cal.Write([]byte(c.Request.URL.RawQuery))
	}
	md5cal.Write([]byte(appSecret.(string)))
	calsign := fmt.Sprintf("%X", md5cal.Sum(nil))
	sign := c.GetHeader("signature")
	if strings.ToUpper(sign) != calsign {
		c.JSON(http.StatusBadRequest, &gin.H{"code": 400001, "msg": "sign error"})
		return false
	}
	return true
}
