package list_test

import (
	"strconv"
	"testing"

	"github.com/smartwalle/stdx/container/list"
)

func BenchmarkList_PushBack_Int(b *testing.B) {
	var l = list.New[int]()

	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
}

func BenchmarkList_PushBack_String(b *testing.B) {
	var l = list.New[string]()

	for i := 0; i < b.N; i++ {
		l.PushBack(strconv.Itoa(i))
	}
}
