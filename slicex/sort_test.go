package slicex_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/smartwalle/stdx/slicex"
)

type Person struct {
	Name  string
	Age   int
	Score float64
}

func cmpIntAsc(a, b int) int {
	return a - b
}

func cmpIntDesc(a, b int) int {
	return b - a
}

func cmpStringAsc(a, b string) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func cmpStringDesc(a, b string) int {
	switch {
	case a > b:
		return -1
	case a < b:
		return 1
	default:
		return 0
	}
}

func cmpFloatDesc(a, b float64) int {
	switch {
	case a > b:
		return -1
	case a < b:
		return 1
	default:
		return 0
	}
}

func TestSortIntSlice(t *testing.T) {
	nums := []int{5, 2, 8, 1, 9}
	slicex.Sort(nums, cmpIntAsc)
	expect := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(nums, expect) {
		t.Errorf("int slice sort failed: expect %v, got %v", expect, nums)
	}
}

func TestSortIntSliceDesc(t *testing.T) {
	nums := []int{5, 2, 8, 1, 9}
	slicex.Sort(nums, cmpIntDesc)
	expect := []int{9, 8, 5, 2, 1}
	if !reflect.DeepEqual(nums, expect) {
		t.Errorf("int slice desc sort failed: expect %v, got %v", expect, nums)
	}
}

func TestSortStringSlice(t *testing.T) {
	strs := []string{"foo", "bar", "baz", "abc"}
	slicex.Sort(strs, cmpStringAsc)
	expect := []string{"abc", "bar", "baz", "foo"}
	if !reflect.DeepEqual(strs, expect) {
		t.Errorf("string slice sort failed: expect %v, got %v", expect, strs)
	}
}

func TestSortStringSliceDesc(t *testing.T) {
	strs := []string{"foo", "bar", "baz", "abc"}
	slicex.Sort(strs, cmpStringDesc)
	expected := []string{"foo", "baz", "bar", "abc"}
	if !reflect.DeepEqual(strs, expected) {
		t.Errorf("string slice desc sort failed: expect %v, got %v", expected, strs)
	}
}

func TestSortStructMultiLevel(t *testing.T) {
	people := []Person{
		{"Alice", 30, 90.5},
		{"Bob", 25, 88.0},
		{"Charlie", 30, 85.0},
		{"Bob", 25, 95.0},
	}
	slicex.Sort(people,
		func(a, b Person) int { return cmpIntAsc(a.Age, b.Age) },
		func(a, b Person) int { return cmpStringAsc(a.Name, b.Name) },
		func(a, b Person) int { return cmpFloatDesc(a.Score, b.Score) },
	)
	expected := []Person{
		{"Bob", 25, 95.0},
		{"Bob", 25, 88.0},
		{"Alice", 30, 90.5},
		{"Charlie", 30, 85.0},
	}
	if !reflect.DeepEqual(people, expected) {
		t.Errorf("struct multi-level sort failed: expect %+v, got %+v", expected, people)
	}
}

func TestSortAllEqual(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{1, 1, 1, 1}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("all equal sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortTime(t *testing.T) {
	n := 1000000
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = Person{
			Name:  fmt.Sprintf("Name%d", rand.Intn(10000)),
			Age:   rand.Intn(100),
			Score: rand.Float64() * 100,
		}
	}
	start := time.Now()
	slicex.Sort(people, func(a, b Person) int { return a.Age - b.Age },
		func(a, b Person) int {
			if a.Score < b.Score {
				return 1
			} else if a.Score > b.Score {
				return -1
			}
			return 0
		},
		func(a, b Person) int {
			if a.Name < b.Name {
				return -1
			} else if a.Name > b.Name {
				return 1
			}
			return 0
		})
	t.Log("耗时(ms):", time.Since(start).Milliseconds())
}

func TestSortEmptySlice(t *testing.T) {
	// 测试空切片排序
	nums := []int{}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("empty slice sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortNilSlice(t *testing.T) {
	// 测试nil切片排序
	var nums []int
	slicex.Sort(nums, cmpIntAsc)
	expected := []int(nil)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("nil slice sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortSingleElement(t *testing.T) {
	// 测试单元素切片排序
	nums := []int{42}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{42}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("single element sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortTwoElements(t *testing.T) {
	// 测试两个元素切片排序
	nums := []int{5, 2}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{2, 5}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("two elements sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortWithNoComparators(t *testing.T) {
	// 测试没有比较函数的情况
	nums := []int{5, 2, 8, 1, 9}
	slicex.Sort(nums)                // 没有传递比较函数
	expected := []int{5, 2, 8, 1, 9} // 应该保持原顺序
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("sort with no comparators failed: expect %v, got %v", expected, nums)
	}
}

func TestSortWithMultipleComparators(t *testing.T) {
	// 测试多个比较函数的情况
	people := []Person{
		{"Alice", 30, 90.5},
		{"Bob", 25, 88.0},
		{"Alice", 30, 85.0},
		{"Bob", 25, 95.0},
	}
	slicex.Sort(people,
		func(a, b Person) int { return cmpIntAsc(a.Age, b.Age) },
		func(a, b Person) int { return cmpStringAsc(a.Name, b.Name) },
	)
	expected := []Person{
		{"Bob", 25, 88.0},
		{"Bob", 25, 95.0},
		{"Alice", 30, 90.5},
		{"Alice", 30, 85.0},
	}
	if !reflect.DeepEqual(people, expected) {
		t.Errorf("sort with multiple comparators failed: expect %+v, got %+v", expected, people)
	}
}

func TestSortWithNegativeNumbers(t *testing.T) {
	// 测试负数排序
	nums := []int{-5, 2, -8, 1, -9, 0}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{-9, -8, -5, 0, 1, 2}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("negative numbers sort failed: expect %v, got %v", expected, nums)
	}
}

func TestSortWithZeroValues(t *testing.T) {
	// 测试包含零值的排序
	nums := []int{0, 1, 0, 2, 0, 3}
	slicex.Sort(nums, cmpIntAsc)
	expected := []int{0, 0, 0, 1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("zero values sort failed: expect %v, got %v", expected, nums)
	}
}
