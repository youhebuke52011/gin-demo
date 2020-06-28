package validater

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

func init() {
	fmt.Println("init validator!!!")
	v := binding.Validator.Engine().(*validator.Validate)
	err := v.RegisterValidation("enum", validEnum)
	fmt.Println(err)
}

