package main

import (
	"testing"
)

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call(&Data{x: 20})
	}
}

func BenchmarkIFaceCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iFaceCall(&Data{x: 20})
	}
}
