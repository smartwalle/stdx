package stdx

import (
	"strconv"
)

func String(value interface{}) string {
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
	default:
		return ""
	}
}
