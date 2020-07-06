package main

import "fmt"

type Tester interface {
	Test(int)
}

type Data struct {
	x int
}

func (d *Data) Test(x int) {
	d.x += x
}

func call(d *Data) {
	d.Test(10)
}

func iFaceCall(t Tester) {
	t.Test(10)
}

func main() {
	d := &Data{x: 10}

	call(d)
	iFaceCall(d)

	fmt.Println(d)
}
