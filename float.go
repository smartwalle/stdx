package stdx

import (
	"strconv"
)

func Float32(value interface{}) float32 {
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

func Float64(value interface{}) float64 {
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
