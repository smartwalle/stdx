package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestProduct(t *testing.T) {
	var l1 = []string{"1", "2", "3", "4"}
	var l2 = []string{"A", "B", "C", "D"}
	var l3 = []string{"★", "☆"}

	var p = [][]string{l1, l2, l3}

	t.Log(slicex.Product(p))
}
