package validater

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"reflect"
)

func validEnum(v validator.FieldLevel) bool {
	value := ""
	enums := strings.Split(v.Param(), "-")
	fvals, fkind, _ := v.ExtractType(v.Field())
	if fkind == reflect.String {
		value = fvals.String()
	} else {
		value = strconv.FormatInt(fvals.Int(), 10)
	}
	for _, enum := range enums {
		if enum == value {
			return true
		}
	}
	return false
}

//func validAPI(v validator.FieldLevel) bool {
//	apis := v.Field().String()
//	for _, r := range Engine.Routes() {
//		if r.Path == apis {
//			return true
//		}
//	}
//	return false
//}

func init() {
	v := binding.Validator.Engine().(*validator.Validate)
	if err := v.RegisterValidation("enum", validEnum); err != nil {
		panic(err)
	}
	log.WithFields(log.Fields{}).Info("validator is ready")
}

