package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(2, 5)

func RateLimit(c *gin.Context) bool {
	if !limiter.Allow() {
		log.WithFields(log.Fields{}).Info("rate limit")
		c.JSON(http.StatusTooManyRequests, &gin.H{})
		return false
	}
	return true
}
