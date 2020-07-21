package singer

import (
	"fmt"
	"gin-demo/client"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func List(c *gin.Context) {
	//data := []map[string]string{
	//	{
	//	"singer_id":   "420",
	//	"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20180622/20180622193316603.jpg",
	//	"singer_name": "陈奕迅",
	//	},
	//	{
	//		"singer_id":   "3520",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20180515/20180515002522714.jpg",
	//		"singer_name": "周杰伦",
	//	},
	//	{
	//		"singer_id":   "3521",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20200706/20200706184106791.jpg",
	//		"singer_name": "张学友",
	//	},
	//	{
	//		"singer_id":   "2351",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20160704/20160704122242573282.jpg",
	//		"singer_name": "齐秦",
	//	},
	//	{
	//		"singer_id":   "1573",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20180507/20180507120242140.jpg",
	//		"singer_name": "刘德华",
	//	},
	//	{
	//		"singer_id":   "3539",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20200116/20200116112014747.jpg",
	//		"singer_name": "张杰",
	//	},
	//	{
	//		"singer_id":   "1574",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20191017/20191017142309922.jpg",
	//		"singer_name": "林俊杰",
	//	},
	//	{
	//		"singer_id":   "93475",
	//		"singer_img":  "http://singerimg.kugou.com/uploadpic/pass/softhead/480/20191209/20191209164452855.jpg",
	//		"singer_name": "李荣浩",
	//	}}
	data := []map[string]interface{}{}
	result := []Singer{}
	fmt.Println(c.Params)
	fmt.Println(c.Request.Form)
	fmt.Println(c.Request.PostForm)
	res := client.GetMysqlCli().Demo.Table("singer")
	if singerId, ok := c.Params.Get("singer_id"); ok {
		res = res.Where("singer_id = ?", singerId)
	}
	res.Find(&result)
	for _, row := range result {
		data = append(data, map[string]interface{}{
			"singer_id":   row.SingerId,
			"singer_name": row.SingerName,
			"singer_img":  row.SingerImg,
		})
	}
	c.JSON(http.StatusOK, &gin.H{"data": gin.H{"singers": data}, "code": 200})
}

func Add(c *gin.Context) {
	var (
		args   = c.Keys["args"].(*AddEntity)
		result = gin.H{"code": 200, "msg": "ok", "data": gin.H{}}
		singer = Singer{
			SingerId:   args.SingerId,
			SingerName: args.SingerName,
			SingerImg:  args.SingerImg,
		}
	)
	if err := client.GetMysqlCli().Demo.Table("singer").Create(&singer).Error; err != nil {
		log.WithFields(log.Fields{"err": err}).Error("gorm")
		result["code"] = 500001
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {
	var (
		args   = c.Keys["args"].(*AddEntity)
		result = gin.H{"code": 200, "msg": "ok", "data": gin.H{}}
		singer Singer
	)
	if err := client.GetMysqlCli().Demo.Table("singer").Where("singer_id = ?", args.SingerId).First(&singer).Error; err != nil {
		log.WithFields(log.Fields{"err": err}).Error("gorm")
		result["code"] = 500001
		c.JSON(http.StatusOK, result)
		return
	}
	singer.SingerName = args.SingerName
	singer.SingerImg = args.SingerImg
	client.GetMysqlCli().Demo.Save(&singer)
	result["data"] = singer.SingerId
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	var singer Singer
	if singerId, ok := c.Params.Get("singer_id"); ok {
		client.GetMysqlCli().Demo.Table("singer").Where("singer_id = ?", singerId).First(&singer)
		client.GetMysqlCli().Demo.Table("singer").Delete(&singer)
	}
	c.JSON(http.StatusOK, &gin.H{"data": singer.SingerId, "code": 200})
}

func Detail(c *gin.Context) {
	var singer Singer
	if singerId, ok := c.Params.Get("singer_id"); ok {
		client.GetMysqlCli().Demo.Table("singer").Where("singer_id = ?", singerId).First(&singer)
	}
	c.JSON(http.StatusOK, &gin.H{
		"data": gin.H{"singer_id": singer.SingerId, "singer_name": singer.SingerName, "singer_img": singer.SingerImg}, "code": 200})
}
