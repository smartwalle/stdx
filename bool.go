package stdx

import (
	"strconv"
)

func Bool(value interface{}) bool {
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
