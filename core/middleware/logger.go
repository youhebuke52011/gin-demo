package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-demo/common"
	"gin-demo/common/response"
	"github.com/gin-gonic/gin"
	"time"
)

type logRowWrite struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

//const SIZE int = 50
var logRowChannel = make(chan string, 100)

func (lgw logRowWrite) Write(sb []byte) (int, error) {
	lgw.buf.Write(sb)
	return lgw.ResponseWriter.Write(sb)
}

func (lgw logRowWrite) WriteString(s string) (int, error) {
	lgw.buf.WriteString(s)
	return lgw.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	go handleAccessLog()
	return func(c *gin.Context) {
		logRow := &logRowWrite{
			ResponseWriter: c.Writer,
			buf:            bytes.NewBufferString(""),
		}

		startTime := time.Now().UnixNano() / 1000
		c.Next()
		resp := response.Resp{}
		if logRow.buf.String() != "" {
			if err := json.Unmarshal(logRow.buf.Bytes(), &resp); err != nil {
				fmt.Printf("logger json:%v\n", err)
			}
		}
		endTime := time.Now().UnixNano() / 1000

		if c.Request.Method == "POST" {
			if err := c.Request.ParseForm(); err != nil {
				fmt.Printf("logger parse: %v\n", err)
			}
		}
		accLogMap := map[string]interface{}{
			"request_time":      startTime,
			"request_method":    c.Request.Method,
			"request_uri":       c.Request.RequestURI,
			"request_proto":     c.Request.Proto,
			"request_ua":        c.Request.UserAgent(),
			"request_referer":   c.Request.Referer(),
			"request_post_data": c.Request.PostForm.Encode(),
			"request_client_ip": c.ClientIP(),
			// ms
			"response_time":     endTime-startTime,
			"response_code":     resp.Code,
			"response_msg":      resp.Msg,
			"response_data":     resp.Data,
		}
		accLogJson, _ := common.Encode(accLogMap)
		logRowChannel <- accLogJson
	}
}

func handleAccessLog() {
	for logrow := range logRowChannel {
		fmt.Println(logrow)
	}
}