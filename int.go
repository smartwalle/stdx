package stdx

import (
	"strconv"
)

func trim(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return s[0:i]
		}
	}
	return s
}

func Int(value interface{}) int {
	switch rValue := value.(type) {
	case int:
		return rValue
	case int8:
		return int(rValue)
	case int16:
		return int(rValue)
	case int32:
		return int(rValue)
	case int64:
		return int(rValue)
	case uint:
		return int(rValue)
	case uint8:
		return int(rValue)
	case uint16:
		return int(rValue)
	case uint32:
		return int(rValue)
	case uint64:
		return int(rValue)
	case uintptr:
		return int(rValue)
	case float32:
		return int(rValue)
	case float64:
		return int(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseInt(trim(rValue), 10, 64)
		return int(nValue)
	default:
		return 0
	}
}

func Int8(value interface{}) int8 {
	switch rValue := value.(type) {
	case int:
		return int8(rValue)
	case int8:
		return rValue
	case int16:
		return int8(rValue)
	case int32:
		return int8(rValue)
	case int64:
		return int8(rValue)
	case uint:
		return int8(rValue)
	case uint8:
		return int8(rValue)
	case uint16:
		return int8(rValue)
	case uint32:
		return int8(rValue)
	case uint64:
		return int8(rValue)
	case uintptr:
		return int8(rValue)
	case float32:
		return int8(rValue)
	case float64:
		return int8(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseInt(trim(rValue), 10, 8)
		return int8(nValue)
	default:
		return 0
	}
}

func Int16(value interface{}) int16 {
	switch rValue := value.(type) {
	case int:
		return int16(rValue)
	case int8:
		return int16(rValue)
	case int16:
		return rValue
	case int32:
		return int16(rValue)
	case int64:
		return int16(rValue)
	case uint:
		return int16(rValue)
	case uint8:
		return int16(rValue)
	case uint16:
		return int16(rValue)
	case uint32:
		return int16(rValue)
	case uint64:
		return int16(rValue)
	case uintptr:
		return int16(rValue)
	case float32:
		return int16(rValue)
	case float64:
		return int16(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseInt(trim(rValue), 10, 16)
		return int16(nValue)
	default:
		return 0
	}
}

func Int32(value interface{}) int32 {
	switch rValue := value.(type) {
	case int:
		return int32(rValue)
	case int8:
		return int32(rValue)
	case int16:
		return int32(rValue)
	case int32:
		return rValue
	case int64:
		return int32(rValue)
	case uint:
		return int32(rValue)
	case uint8:
		return int32(rValue)
	case uint16:
		return int32(rValue)
	case uint32:
		return int32(rValue)
	case uint64:
		return int32(rValue)
	case uintptr:
		return int32(rValue)
	case float32:
		return int32(rValue)
	case float64:
		return int32(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseInt(trim(rValue), 10, 32)
		return int32(nValue)
	default:
		return 0
	}
}

func Int64(value interface{}) int64 {
	switch rValue := value.(type) {
	case int:
		return int64(rValue)
	case int8:
		return int64(rValue)
	case int16:
		return int64(rValue)
	case int32:
		return int64(rValue)
	case int64:
		return rValue
	case uint:
		return int64(rValue)
	case uint8:
		return int64(rValue)
	case uint16:
		return int64(rValue)
	case uint32:
		return int64(rValue)
	case uint64:
		return int64(rValue)
	case uintptr:
		return int64(rValue)
	case float32:
		return int64(rValue)
	case float64:
		return int64(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseInt(trim(rValue), 10, 64)
		return nValue
	default:
		return 0
	}
}

func Uint(value interface{}) uint {
	switch rValue := value.(type) {
	case int:
		return uint(rValue)
	case int8:
		return uint(rValue)
	case int16:
		return uint(rValue)
	case int32:
		return uint(rValue)
	case int64:
		return uint(rValue)
	case uint:
		return rValue
	case uint8:
		return uint(rValue)
	case uint16:
		return uint(rValue)
	case uint32:
		return uint(rValue)
	case uint64:
		return uint(rValue)
	case uintptr:
		return uint(rValue)
	case float32:
		return uint(rValue)
	case float64:
		return uint(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 64)
		return uint(nValue)
	default:
		return 0
	}
}

func Uint8(value interface{}) uint8 {
	switch rValue := value.(type) {
	case int:
		return uint8(rValue)
	case int8:
		return uint8(rValue)
	case int16:
		return uint8(rValue)
	case int32:
		return uint8(rValue)
	case int64:
		return uint8(rValue)
	case uint:
		return uint8(rValue)
	case uint8:
		return rValue
	case uint16:
		return uint8(rValue)
	case uint32:
		return uint8(rValue)
	case uint64:
		return uint8(rValue)
	case uintptr:
		return uint8(rValue)
	case float32:
		return uint8(rValue)
	case float64:
		return uint8(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 8)
		return uint8(nValue)
	default:
		return 0
	}
}

func Uint16(value interface{}) uint16 {
	switch rValue := value.(type) {
	case int:
		return uint16(rValue)
	case int8:
		return uint16(rValue)
	case int16:
		return uint16(rValue)
	case int32:
		return uint16(rValue)
	case int64:
		return uint16(rValue)
	case uint:
		return uint16(rValue)
	case uint8:
		return uint16(rValue)
	case uint16:
		return rValue
	case uint32:
		return uint16(rValue)
	case uint64:
		return uint16(rValue)
	case uintptr:
		return uint16(rValue)
	case float32:
		return uint16(rValue)
	case float64:
		return uint16(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 16)
		return uint16(nValue)
	default:
		return 0
	}
}

func Uint32(value interface{}) uint32 {
	switch rValue := value.(type) {
	case int:
		return uint32(rValue)
	case int8:
		return uint32(rValue)
	case int16:
		return uint32(rValue)
	case int32:
		return uint32(rValue)
	case int64:
		return uint32(rValue)
	case uint:
		return uint32(rValue)
	case uint8:
		return uint32(rValue)
	case uint16:
		return uint32(rValue)
	case uint32:
		return rValue
	case uint64:
		return uint32(rValue)
	case uintptr:
		return uint32(rValue)
	case float32:
		return uint32(rValue)
	case float64:
		return uint32(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 32)
		return uint32(nValue)
	default:
		return 0
	}
}

func Uint64(value interface{}) uint64 {
	switch rValue := value.(type) {
	case int:
		return uint64(rValue)
	case int8:
		return uint64(rValue)
	case int16:
		return uint64(rValue)
	case int32:
		return uint64(rValue)
	case int64:
		return uint64(rValue)
	case uint:
		return uint64(rValue)
	case uint8:
		return uint64(rValue)
	case uint16:
		return uint64(rValue)
	case uint32:
		return uint64(rValue)
	case uint64:
		return rValue
	case uintptr:
		return uint64(rValue)
	case float32:
		return uint64(rValue)
	case float64:
		return uint64(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 64)
		return nValue
	default:
		return 0
	}
}

func Uintptr(value interface{}) uintptr {
	switch rValue := value.(type) {
	case int:
		return uintptr(rValue)
	case int8:
		return uintptr(rValue)
	case int16:
		return uintptr(rValue)
	case int32:
		return uintptr(rValue)
	case int64:
		return uintptr(rValue)
	case uint:
		return uintptr(rValue)
	case uint8:
		return uintptr(rValue)
	case uint16:
		return uintptr(rValue)
	case uint32:
		return uintptr(rValue)
	case uint64:
		return uintptr(rValue)
	case uintptr:
		return rValue
	case float32:
		return uintptr(rValue)
	case float64:
		return uintptr(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseUint(trim(rValue), 10, 64)
		return uintptr(nValue)
	default:
		return 0
	}
}
