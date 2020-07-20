package tg

import (
	"encoding/json"
	"gin-demo/client"
	"gin-demo/common/compression"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func TGzip(ctx *gin.Context) {
	var (
		rds = client.GetRedisCli().Demo
	)
	if resTmp, err := rds.Get(ctx, "gz").Result(); err == nil {
		res := gin.H{}
		err := json.Unmarshal([]byte(resTmp), &res)
		if err != nil {
			log.WithFields(log.Fields{"err": err.Error()}).Error("json")
		}
		ctx.JSON(http.StatusOK, &gin.H{"data": res})
		return
	}
	res := map[string]interface{}{}
	compression.KGDecodeAndGz(tmp, &res)
	sb, err := json.Marshal(res)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error("json")
	}
	if err := rds.Set(ctx, "gz", string(sb), 1*time.Second).Err(); err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error("redis")
	}
	ctx.JSON(http.StatusOK, &gin.H{"data": res})
}