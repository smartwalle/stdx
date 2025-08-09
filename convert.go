package stdx

import (
	"strconv"
	"strings"
	"unsafe"
)

func ToBool(value interface{}) bool {
	switch rValue := value.(type) {
	case int:
		return rValue != 0
	case int8:
		return rValue != 0
	case int16:
		return rValue != 0
	case int32:
		return rValue != 0
	case int64:
		return rValue != 0
	case uint:
		return rValue != 0
	case uint8:
		return rValue != 0
	case uint16:
		return rValue != 0
	case uint32:
		return rValue != 0
	case uint64:
		return rValue != 0
	case uintptr:
		return rValue != 0
	case float32:
		return rValue != 0
	case float64:
		return rValue != 0
	case bool:
		return rValue
	case string:
		var nValue, _ = strconv.ParseBool(rValue)
		return nValue
	default:
		return false
	}
}

func ToFloat32(value interface{}) float32 {
	switch rValue := value.(type) {
	case int:
		return float32(rValue)
	case int8:
		return float32(rValue)
	case int16:
		return float32(rValue)
	case int32:
		return float32(rValue)
	case int64:
		return float32(rValue)
	case uint:
		return float32(rValue)
	case uint8:
		return float32(rValue)
	case uint16:
		return float32(rValue)
	case uint32:
		return float32(rValue)
	case uint64:
		return float32(rValue)
	case uintptr:
		return float32(rValue)
	case float32:
		return rValue
	case float64:
		return float32(rValue)
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseFloat(rValue, 32)
		return float32(nValue)
	default:
		return 0
	}
}

func ToFloat64(value interface{}) float64 {
	switch rValue := value.(type) {
	case int:
		return float64(rValue)
	case int8:
		return float64(rValue)
	case int16:
		return float64(rValue)
	case int32:
		return float64(rValue)
	case int64:
		return float64(rValue)
	case uint:
		return float64(rValue)
	case uint8:
		return float64(rValue)
	case uint16:
		return float64(rValue)
	case uint32:
		return float64(rValue)
	case uint64:
		return float64(rValue)
	case uintptr:
		return float64(rValue)
	case float32:
		return float64(rValue)
	case float64:
		return rValue
	case bool:
		if rValue {
			return 1
		}
		return 0
	case string:
		var nValue, _ = strconv.ParseFloat(rValue, 64)
		return nValue
	default:
		return 0
	}
}

func ToInt(value interface{}) int {
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
		var nValue, _ = strconv.ParseInt(trimDecimal(rValue), 10, 64)
		return int(nValue)
	default:
		return 0
	}
}

func ToInt8(value interface{}) int8 {
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
		var nValue, _ = strconv.ParseInt(trimDecimal(rValue), 10, 8)
		return int8(nValue)
	default:
		return 0
	}
}

func ToInt16(value interface{}) int16 {
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
		var nValue, _ = strconv.ParseInt(trimDecimal(rValue), 10, 16)
		return int16(nValue)
	default:
		return 0
	}
}

func ToInt32(value interface{}) int32 {
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
		var nValue, _ = strconv.ParseInt(trimDecimal(rValue), 10, 32)
		return int32(nValue)
	default:
		return 0
	}
}

func ToInt64(value interface{}) int64 {
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
		var nValue, _ = strconv.ParseInt(trimDecimal(rValue), 10, 64)
		return nValue
	default:
		return 0
	}
}

func ToUint(value interface{}) uint {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return uint(nValue)
	default:
		return 0
	}
}

func ToUint8(value interface{}) uint8 {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 8)
		return uint8(nValue)
	default:
		return 0
	}
}

func ToUint16(value interface{}) uint16 {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 16)
		return uint16(nValue)
	default:
		return 0
	}
}

func ToUint32(value interface{}) uint32 {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 32)
		return uint32(nValue)
	default:
		return 0
	}
}

func ToUint64(value interface{}) uint64 {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return nValue
	default:
		return 0
	}
}

func ToUintptr(value interface{}) uintptr {
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
		var nValue, _ = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return uintptr(nValue)
	default:
		return 0
	}
}

func ToString(value interface{}) string {
	switch rValue := value.(type) {
	case int:
		return strconv.FormatInt(int64(rValue), 10)
	case int8:
		return strconv.FormatInt(int64(rValue), 10)
	case int16:
		return strconv.FormatInt(int64(rValue), 10)
	case int32:
		return strconv.FormatInt(int64(rValue), 10)
	case int64:
		return strconv.FormatInt(rValue, 10)
	case uint:
		return strconv.FormatUint(uint64(rValue), 10)
	case uint8:
		return strconv.FormatUint(uint64(rValue), 10)
	case uint16:
		return strconv.FormatUint(uint64(rValue), 10)
	case uint32:
		return strconv.FormatUint(uint64(rValue), 10)
	case uint64:
		return strconv.FormatUint(rValue, 10)
	case uintptr:
		return strconv.FormatUint(uint64(rValue), 10)
	case float32:
		return strconv.FormatFloat(float64(rValue), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(rValue, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(rValue)
	case string:
		return rValue
	case []byte:
		return unsafe.String(&rValue[0], len(rValue))
	case []rune:
		return string(rValue)
	default:
		return ""
	}
}

func trimDecimal(s string) string {
	if idx := strings.IndexByte(s, '.'); idx != -1 {
		return s[:idx]
	}
	return s
}
