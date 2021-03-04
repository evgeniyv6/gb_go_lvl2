package main

import "testing"

var (
	GlobalRes int
)

func BenchmarkKiloGoroutins(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = kiloGoroutins()
	}
	GlobalRes = res
}
