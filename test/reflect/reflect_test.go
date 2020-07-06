package main

import (
	"testing"
)

func BenchmarkIncX(b *testing.B) {
	d := struct {
		X int
	}{10}
	for i := 0; i < b.N; i++ {
		incX(&d)
	}
}

func BenchmarkUnsafeIncX(b *testing.B) {
	d := struct {
		X int
	}{10}
	for i := 0; i < b.N; i++ {
		unsafeIncx(&d)
	}
}

func BenchmarkUseMapIncX(b *testing.B) {
	d := struct {
		X int
	}{10}
	for i := 0; i < b.N; i++ {
		unsafeUseMapIncx(&d)
	}
}