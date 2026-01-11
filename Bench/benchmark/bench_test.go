package main

import "testing"

var t []int

func BenchmarkSliceNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceNew(i)
	}
}
func BenchmarkSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceAppend(i)
	}
}
