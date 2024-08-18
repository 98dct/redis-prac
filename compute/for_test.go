package main

import "testing"

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forRange()
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursion()
	}
}
