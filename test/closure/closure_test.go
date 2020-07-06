package main

import "testing"

func test(x int) int {
	return x * 2
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test(i)
	}
}

func BenchmarkClosureTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func() int {
			return i * 2
		}
	}
}

func BenchmarkAnonymousTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func(k int) int {
			return k * 2
		}(i)
	}
}