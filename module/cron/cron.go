package cron

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var C *cron.Cron

func Close() {
	log.WithFields(log.Fields{}).Info("close")
	C.Stop()
}

func init() {
	C = cron.New(cron.WithSeconds())
	C.AddFunc("* 30 * * * *", func() {
		log.WithFields(log.Fields{}).Info("test")
		//fmt.Println("Every hour on the half hour")
	})
	C.Start()
	log.WithFields(log.Fields{}).Info("cron is ready")
}
