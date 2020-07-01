package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var C *cron.Cron

func Close() {
	C.Stop()
}

func init() {
	C = cron.New(cron.WithSeconds())
	C.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	C.Start()
	log.WithFields(log.Fields{}).Info("cron is ready")
}
