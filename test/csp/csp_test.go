package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func chanCounter() chan int {
	c := make(chan int)
	go func() {
		for i := 1; ; i++ {
			c <- i
		}
	}()
	return c
}

func mutexCounter() func() int {
	var (
		m sync.Mutex
		x int
	)
	return func() (res int) {
		m.Lock()
		x++
		res = x
		m.Unlock()
		return
	}
}

func atomicCounter() func() int {
	var x int64
	return func() int {
		return int(atomic.AddInt64(&x, 1))
	}
}

func BenchmarkChanCounter(b *testing.B) {
	cc := chanCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = <-cc
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	mc := mutexCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mc()
	}
}

func BenchmarkAtomicCounter(b *testing.B) {
	ac := atomicCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ac()
	}
}