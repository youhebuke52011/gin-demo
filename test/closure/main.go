package main

import "time"

func closureTest() {
	x := 100
	go func() {
		x++
		println(x)
	}()
	x++
}

func test() {
	y := 100
	go func(k int) {
		k++
		println(k)
	}(y)
	y++
}

func main() {
	closureTest()
	test()
	time.Sleep(2 * time.Second)
}
