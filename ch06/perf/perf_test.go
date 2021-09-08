package perf

import "testing"

func BenchmarkOne(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N;i++ {
		One()
	}
}

func BenchmarkTwo(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		Two()
	}
}

//to run
//go test -v -bench=. -benchmem