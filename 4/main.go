package main

import (
	"strings"
	"testing"
)

func BenchmarkDefaultConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 1000; j++ {
			s += "x"
		}
	}
}

func BenchmarkConcatWithStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 1000; j++ {
			sb.WriteString("x")
		}
		_ = sb.String()
	}
}
