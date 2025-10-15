package mapx

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
)

type OrderedMap[K Key, V any] struct {
	keys   []K
	values map[K]V
}

func NewOrderedMap[K Key, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

func (m *OrderedMap[K, V]) Range(fn func(key K, value V) bool) {
	if len(m.keys) == 0 || len(m.values) == 0 {
		return
	}
	for _, key := range m.keys {
		if ok := fn(key, m.values[key]); !ok {
			return
		}
	}
}

func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	if len(m.values) == 0 {
		return value, false
	}
	value, exists := m.values[key]
	return value, exists
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	if m.values == nil {
		m.values = make(map[K]V)
	}
	if _, exists := m.values[key]; !exists {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *OrderedMap[K, V]) Delete(key K) {
	if len(m.values) == 0 {
		return
	}

	if _, exists := m.values[key]; !exists {
		return
	}
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	delete(m.values, key)
}

func (m *OrderedMap[K, V]) Keys() []K {
	var keys = make([]K, len(m.keys))
	for idx, key := range m.keys {
		keys[idx] = key
	}
	return keys
}

func (m *OrderedMap[K, V]) Values() []V {
	var values = make([]V, len(m.keys))
	for idx, key := range m.keys {
		values[idx] = m.values[key]
	}
	return values
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.values)
}

func (m *OrderedMap[K, V]) Clear() {
	m.keys = make([]K, 0)
	m.values = make(map[K]V)
}

func (m *OrderedMap[K, V]) UnmarshalJSON(b []byte) error {
	m.keys = make([]K, 0)
	m.values = make(map[K]V)
	if err := m.decode(b); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (m *OrderedMap[K, V]) decode(b []byte) error {
	var rMap = make(map[K]V)
	if err := json.Unmarshal(b, &rMap); err != nil {
		return err
	}

	var decoder = json.NewDecoder(bytes.NewReader(b))
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case '{':
				continue
			case '[':
				continue
			case '}':
				continue
			case ']':
				continue
			default:
			}
		}

		var rawKey = token.(string)

		var key K
		switch any(key).(type) {
		case string:
			key = any(rawKey).(K)
		case int:
			intValue, err := strconv.ParseInt(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(int(intValue)).(K)
		case int8:
			intValue, err := strconv.ParseInt(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(int8(intValue)).(K)
		case int16:
			intValue, err := strconv.ParseInt(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(int16(intValue)).(K)
		case int32:
			intValue, err := strconv.ParseInt(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(int32(intValue)).(K)
		case int64:
			intValue, err := strconv.ParseInt(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(intValue).(K)
		case uint:
			intValue, err := strconv.ParseUint(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(uint(intValue)).(K)
		case uint8:
			intValue, err := strconv.ParseUint(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(uint8(intValue)).(K)
		case uint16:
			intValue, err := strconv.ParseUint(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(uint16(intValue)).(K)
		case uint32:
			intValue, err := strconv.ParseUint(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(uint32(intValue)).(K)
		case uint64:
			intValue, err := strconv.ParseUint(rawKey, 10, 64)
			if err != nil {
				return err
			}
			key = any(intValue).(K)
		}

		if value, exists := rMap[key]; exists {
			m.Set(key, value)
		}

		token, err = decoder.Token()
		if err != nil {
			return err
		}

		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case '{':
				if err = skip(decoder, '}'); err != nil {
					return err
				}
			case '[':
				if err = skip(decoder, ']'); err != nil {
					return err
				}
			default:
			}
		}
	}
}

func skip(decoder *json.Decoder, s json.Delim) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case '{':
				return skip(decoder, '}')
			case '[':
				return skip(decoder, ']')
			case s:
				return nil
			default:
			}
		}
	}
}

func (m OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	if m.keys == nil || m.values == nil {
		return []byte("null"), nil
	}

	var buf = bytes.NewBufferString("{")
	var encoder = json.NewEncoder(buf)

	for idx, key := range m.keys {
		if idx > 0 {
			if err := buf.WriteByte(','); err != nil {
				return nil, err
			}
		}

		switch rKey := any(key).(type) {
		case string:
			if err := encoder.Encode(rKey); err != nil {
				return nil, err
			}
		case int:
			if err := encoder.Encode(strconv.FormatInt(int64(rKey), 10)); err != nil {
				return nil, err
			}
		case int8:
			if err := encoder.Encode(strconv.FormatInt(int64(rKey), 10)); err != nil {
				return nil, err
			}
		case int16:
			if err := encoder.Encode(strconv.FormatInt(int64(rKey), 10)); err != nil {
				return nil, err
			}
		case int32:
			if err := encoder.Encode(strconv.FormatInt(int64(rKey), 10)); err != nil {
				return nil, err
			}
		case int64:
			if err := encoder.Encode(strconv.FormatInt(rKey, 10)); err != nil {
				return nil, err
			}
		case uint:
			if err := encoder.Encode(strconv.FormatUint(uint64(rKey), 10)); err != nil {
				return nil, err
			}
		case uint8:
			if err := encoder.Encode(strconv.FormatUint(uint64(rKey), 10)); err != nil {
				return nil, err
			}
		case uint16:
			if err := encoder.Encode(strconv.FormatUint(uint64(rKey), 10)); err != nil {
				return nil, err
			}
		case uint32:
			if err := encoder.Encode(strconv.FormatUint(uint64(rKey), 10)); err != nil {
				return nil, err
			}
		case uint64:
			if err := encoder.Encode(strconv.FormatUint(rKey, 10)); err != nil {
				return nil, err
			}
		default:
			keyBytes, err := json.Marshal(key)
			if err != nil {
				return nil, err
			}
			if err = encoder.Encode(string(keyBytes)); err != nil {
				return nil, err
			}
		}

		if err := buf.WriteByte(':'); err != nil {
			return nil, err
		}
		if err := encoder.Encode(m.values[key]); err != nil {
			return nil, err
		}
	}

	if err := buf.WriteByte('}'); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
