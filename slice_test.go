package stdx_test

import (
	"github.com/smartwalle/stdx"
	"testing"
)

func BenchmarkStdAppend(b *testing.B) {
	var ints []int
	for i := 0; i < b.N; i++ {
		ints = append(ints, i)
	}
}

func BenchmarkAppend(b *testing.B) {
	var ints []int
	for i := 0; i < b.N; i++ {
		ints = stdx.Append(ints, i)
	}
}
