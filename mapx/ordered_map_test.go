package mapx_test

import (
	"encoding/json"
	"github.com/smartwalle/stdx/mapx"
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	var m1 = mapx.NewOrderedMap[string, string]()
	m1.Set("1", "v1")
	m1.Set("11", "v11")
	m1.Set("2", "v2")
	m1.Set("22", "v22")
	m1.Set("4", "v4")
	m1.Set("44", "v44")
	m1.Set("3", "v3")
	m1.Set("33", "v33")

	check(t, m1, []string{"1", "11", "2", "22", "4", "44", "3", "33"}, `{"1":"v1","11":"v11","2":"v2","22":"v22","4":"v4","44":"v44","3":"v3","33":"v33"}`)

	m1.Delete("4")
	check(t, m1, []string{"1", "11", "2", "22", "44", "3", "33"}, `{"1":"v1","11":"v11","2":"v2","22":"v22","44":"v44","3":"v3","33":"v33"}`)

	m1.Set("4", "v4")
	check(t, m1, []string{"1", "11", "2", "22", "44", "3", "33", "4"}, `{"1":"v1","11":"v11","2":"v2","22":"v22","44":"v44","3":"v3","33":"v33","4":"v4"}`)

	m1.Set("中文", "中文内容")
	check(t, m1, []string{"1", "11", "2", "22", "44", "3", "33", "4", "中文"}, `{"1":"v1","11":"v11","2":"v2","22":"v22","44":"v44","3":"v3","33":"v33","4":"v4","中文":"中文内容"}`)

	var m2 *mapx.OrderedMap[int, string]
	jsonBytes, _ := json.Marshal(m2)
	if string(jsonBytes) != "null" {
		t.Fatal("JSON序列化异常")
	}
	var m3 mapx.OrderedMap[int, string]
	jsonBytes, _ = json.Marshal(m3)
	if string(jsonBytes) != "{}" {
		t.Fatal("JSON序列化异常")
	}
}

func check[K mapx.Key, V any](t *testing.T, m *mapx.OrderedMap[K, V], keys []K, jsonString string) {
	// 验证Key顺序
	if !slicex.Equals(m.Keys(), keys, func(a, b K) bool {
		return a == b
	}) {
		t.Fatalf("顺序异常, 期望: %+v, 实际: %+v", keys, m.Keys())
	}

	// 验证Range顺序
	var idx = 0
	m.Range(func(key K, value V) bool {
		if key != keys[idx] {
			t.Fatalf("顺序异常, 期望: %+v, 实际: %+v", keys[idx], key)
		}
		idx++
		return true
	})

	// 验证JSON序列化顺序
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if string(jsonBytes) != jsonString {
		t.Fatalf("JSON序列化异常, 期望: %s, 实际: %s", jsonString, string(jsonBytes))
	}

	// 验证JSON反序列化之后的顺序
	var m2 mapx.OrderedMap[K, V]
	if err = json.Unmarshal(jsonBytes, &m2); err != nil {
		t.Fatal(err)
	}
	if !slicex.Equals(m.Keys(), m2.Keys(), func(a, b K) bool {
		return a == b
	}) {
		t.Fatalf("JSON反序列化之后顺序异常, 期望: %+v, 实际: %+v", m.Keys(), m2.Keys())
	}

	var m3 *mapx.OrderedMap[K, V]
	if err = json.Unmarshal(jsonBytes, &m3); err != nil {
		t.Fatal(err)
	}
	if !slicex.Equals(m.Keys(), m3.Keys(), func(a, b K) bool {
		return a == b
	}) {
		t.Fatalf("JSON反序列化之后顺序异常, 期望: %+v, 实际: %+v", m.Keys(), m2.Keys())
	}
}
