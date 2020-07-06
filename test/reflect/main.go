package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func incX(d interface{}) int64 {
	//v := reflect.ValueOf(d).Elem()
	v := reflect.Indirect(reflect.ValueOf(d))
	f := v.FieldByName("X")

	i := f.Int()
	i++
	f.SetInt(i)
	return i
}

var offset uintptr = 0xFFFF

func unsafeIncx(d interface{}) int64 {
	if offset == 0xFFFF {
		//v := reflect.Indirect(reflect.ValueOf(d))
		v := reflect.TypeOf(d).Elem()
		f, _ := v.FieldByName("X")
		offset = f.Offset
	}
	p := (*[2]uintptr)(unsafe.Pointer(&d))
	px := (*int64)(unsafe.Pointer(p[1] + offset))
	*px++
	return *px
}

var cache = map[*uintptr]map[string]uintptr{}

func unsafeUseMapIncx(d interface{}) int64 {
	itab := *(**uintptr)(unsafe.Pointer(&d))
	m, ok := cache[itab]
	if !ok {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	offset, ok := m["X"]
	if !ok {
		v := reflect.TypeOf(d).Elem()
		f, _ := v.FieldByName("X")
		offset = f.Offset
		m["X"] = offset
	}
	p := (*[2]uintptr)(unsafe.Pointer(&d))
	px := (*int64)(unsafe.Pointer(p[1] + offset))
	*px++
	return *px
}

func main() {
	d := struct {
		X int
	}{10}
	//res := incX(&d)
	res := unsafeIncx(&d)
	fmt.Println(res)
	fmt.Println(d.X)
}
