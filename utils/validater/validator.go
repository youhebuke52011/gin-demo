package validater

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	//"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

//func validEnum(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
func ValidEnum(fl validator.FieldLevel) bool {
	//if fl.Field().String() == "invalid" {
	//	return false
	//}
	fmt.Printf("valid enum:%v\n", fl.Field().String())
	fmt.Printf("valid enum:%v\n", fl.Field().Int())
	return false
	//value := ""
	//enums := strings.Split(param, "-")
	//if fieldKind == reflect.String {
	//	value = field.String()
	//} else {
	//	value = strconv.FormatInt(field.Int(), 10)
	//}
	//for _, enum := range enums {
	//	if enum == value {
	//		return true
	//	}
	//}
	//return false
}

func validTimestamp(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	tims := field.Int()
	nows := time.Now().Unix()
	if nows-tims >= 24*3600 || tims-nows >= 24*3600 {
		return false
	}
	return true
}

func validDatetime(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	datetime := field.String()
	if _, err := time.Parse("", datetime); err != nil {
		return false
	}
	return true
}
//func customFunc(fl validator.FieldLevel) bool {
//
//	if fl.Field().String() == "invalid" {
//		return false
//	}
//
//	return true
//}
//
//validate.RegisterValidation("custom tag name", customFunc)

func init() {
	fmt.Println("init validator!!!")
	v := binding.Validator.Engine().(*validator.Validate)
	err := v.RegisterValidation("enum", ValidEnum)
	fmt.Println(err)
	//v.RegisterValidation("enum", validEnum)
	//v.RegisterValidation("timestamp", validTimestamp)
	//v.RegisterValidation("datetime", validDatetime)
}

