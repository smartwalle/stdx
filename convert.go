package stdx

import (
	"errors"
	"strconv"
	"strings"
)

var ErrUnsupportedValue = errors.New("unsupported value")

func MustBool(value interface{}, defaultValue bool) bool {
	var rValue, err = ToBool(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToBool(value interface{}) (bool, error) {
	switch rValue := value.(type) {
	case int:
		return rValue != 0, nil
	case int8:
		return rValue != 0, nil
	case int16:
		return rValue != 0, nil
	case int32:
		return rValue != 0, nil
	case int64:
		return rValue != 0, nil
	case uint:
		return rValue != 0, nil
	case uint8:
		return rValue != 0, nil
	case uint16:
		return rValue != 0, nil
	case uint32:
		return rValue != 0, nil
	case uint64:
		return rValue != 0, nil
	case uintptr:
		return rValue != 0, nil
	case float32:
		return rValue != 0, nil
	case float64:
		return rValue != 0, nil
	case bool:
		return rValue, nil
	case string:
		var nValue, err = strconv.ParseBool(rValue)
		return nValue, err
	default:
		return false, ErrUnsupportedValue
	}
}

func MustFloat32(value interface{}, defaultValue float32) float32 {
	var rValue, err = ToFloat32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToFloat32(value interface{}) (float32, error) {
	switch rValue := value.(type) {
	case int:
		return float32(rValue), nil
	case int8:
		return float32(rValue), nil
	case int16:
		return float32(rValue), nil
	case int32:
		return float32(rValue), nil
	case int64:
		return float32(rValue), nil
	case uint:
		return float32(rValue), nil
	case uint8:
		return float32(rValue), nil
	case uint16:
		return float32(rValue), nil
	case uint32:
		return float32(rValue), nil
	case uint64:
		return float32(rValue), nil
	case uintptr:
		return float32(rValue), nil
	case float32:
		return rValue, nil
	case float64:
		return float32(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseFloat(rValue, 32)
		return float32(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustFloat64(value interface{}, defaultValue float64) float64 {
	var rValue, err = ToFloat64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToFloat64(value interface{}) (float64, error) {
	switch rValue := value.(type) {
	case int:
		return float64(rValue), nil
	case int8:
		return float64(rValue), nil
	case int16:
		return float64(rValue), nil
	case int32:
		return float64(rValue), nil
	case int64:
		return float64(rValue), nil
	case uint:
		return float64(rValue), nil
	case uint8:
		return float64(rValue), nil
	case uint16:
		return float64(rValue), nil
	case uint32:
		return float64(rValue), nil
	case uint64:
		return float64(rValue), nil
	case uintptr:
		return float64(rValue), nil
	case float32:
		return float64(rValue), nil
	case float64:
		return rValue, nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseFloat(rValue, 64)
		return nValue, err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustInt(value interface{}, defaultValue int) int {
	var rValue, err = ToInt(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToInt(value interface{}) (int, error) {
	switch rValue := value.(type) {
	case int:
		return rValue, nil
	case int8:
		return int(rValue), nil
	case int16:
		return int(rValue), nil
	case int32:
		return int(rValue), nil
	case int64:
		return int(rValue), nil
	case uint:
		return int(rValue), nil
	case uint8:
		return int(rValue), nil
	case uint16:
		return int(rValue), nil
	case uint32:
		return int(rValue), nil
	case uint64:
		return int(rValue), nil
	case uintptr:
		return int(rValue), nil
	case float32:
		return int(rValue), nil
	case float64:
		return int(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseInt(trimDecimal(rValue), 10, 64)
		return int(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustInt8(value interface{}, defaultValue int8) int8 {
	var rValue, err = ToInt8(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToInt8(value interface{}) (int8, error) {
	switch rValue := value.(type) {
	case int:
		return int8(rValue), nil
	case int8:
		return rValue, nil
	case int16:
		return int8(rValue), nil
	case int32:
		return int8(rValue), nil
	case int64:
		return int8(rValue), nil
	case uint:
		return int8(rValue), nil
	case uint8:
		return int8(rValue), nil
	case uint16:
		return int8(rValue), nil
	case uint32:
		return int8(rValue), nil
	case uint64:
		return int8(rValue), nil
	case uintptr:
		return int8(rValue), nil
	case float32:
		return int8(rValue), nil
	case float64:
		return int8(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseInt(trimDecimal(rValue), 10, 8)
		return int8(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustInt16(value interface{}, defaultValue int16) int16 {
	var rValue, err = ToInt16(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToInt16(value interface{}) (int16, error) {
	switch rValue := value.(type) {
	case int:
		return int16(rValue), nil
	case int8:
		return int16(rValue), nil
	case int16:
		return rValue, nil
	case int32:
		return int16(rValue), nil
	case int64:
		return int16(rValue), nil
	case uint:
		return int16(rValue), nil
	case uint8:
		return int16(rValue), nil
	case uint16:
		return int16(rValue), nil
	case uint32:
		return int16(rValue), nil
	case uint64:
		return int16(rValue), nil
	case uintptr:
		return int16(rValue), nil
	case float32:
		return int16(rValue), nil
	case float64:
		return int16(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseInt(trimDecimal(rValue), 10, 16)
		return int16(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustInt32(value interface{}, defaultValue int32) int32 {
	var rValue, err = ToInt32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToInt32(value interface{}) (int32, error) {
	switch rValue := value.(type) {
	case int:
		return int32(rValue), nil
	case int8:
		return int32(rValue), nil
	case int16:
		return int32(rValue), nil
	case int32:
		return rValue, nil
	case int64:
		return int32(rValue), nil
	case uint:
		return int32(rValue), nil
	case uint8:
		return int32(rValue), nil
	case uint16:
		return int32(rValue), nil
	case uint32:
		return int32(rValue), nil
	case uint64:
		return int32(rValue), nil
	case uintptr:
		return int32(rValue), nil
	case float32:
		return int32(rValue), nil
	case float64:
		return int32(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseInt(trimDecimal(rValue), 10, 32)
		return int32(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustInt64(value interface{}, defaultValue int64) int64 {
	var rValue, err = ToInt64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToInt64(value interface{}) (int64, error) {
	switch rValue := value.(type) {
	case int:
		return int64(rValue), nil
	case int8:
		return int64(rValue), nil
	case int16:
		return int64(rValue), nil
	case int32:
		return int64(rValue), nil
	case int64:
		return rValue, nil
	case uint:
		return int64(rValue), nil
	case uint8:
		return int64(rValue), nil
	case uint16:
		return int64(rValue), nil
	case uint32:
		return int64(rValue), nil
	case uint64:
		return int64(rValue), nil
	case uintptr:
		return int64(rValue), nil
	case float32:
		return int64(rValue), nil
	case float64:
		return int64(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseInt(trimDecimal(rValue), 10, 64)
		return nValue, err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUint(value interface{}, defaultValue uint) uint {
	var rValue, err = ToUint(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUint(value interface{}) (uint, error) {
	switch rValue := value.(type) {
	case int:
		return uint(rValue), nil
	case int8:
		return uint(rValue), nil
	case int16:
		return uint(rValue), nil
	case int32:
		return uint(rValue), nil
	case int64:
		return uint(rValue), nil
	case uint:
		return rValue, nil
	case uint8:
		return uint(rValue), nil
	case uint16:
		return uint(rValue), nil
	case uint32:
		return uint(rValue), nil
	case uint64:
		return uint(rValue), nil
	case uintptr:
		return uint(rValue), nil
	case float32:
		return uint(rValue), nil
	case float64:
		return uint(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return uint(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUint8(value interface{}, defaultValue uint8) uint8 {
	var rValue, err = ToUint8(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUint8(value interface{}) (uint8, error) {
	switch rValue := value.(type) {
	case int:
		return uint8(rValue), nil
	case int8:
		return uint8(rValue), nil
	case int16:
		return uint8(rValue), nil
	case int32:
		return uint8(rValue), nil
	case int64:
		return uint8(rValue), nil
	case uint:
		return uint8(rValue), nil
	case uint8:
		return rValue, nil
	case uint16:
		return uint8(rValue), nil
	case uint32:
		return uint8(rValue), nil
	case uint64:
		return uint8(rValue), nil
	case uintptr:
		return uint8(rValue), nil
	case float32:
		return uint8(rValue), nil
	case float64:
		return uint8(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 8)
		return uint8(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUint16(value interface{}, defaultValue uint16) uint16 {
	var rValue, err = ToUint16(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUint16(value interface{}) (uint16, error) {
	switch rValue := value.(type) {
	case int:
		return uint16(rValue), nil
	case int8:
		return uint16(rValue), nil
	case int16:
		return uint16(rValue), nil
	case int32:
		return uint16(rValue), nil
	case int64:
		return uint16(rValue), nil
	case uint:
		return uint16(rValue), nil
	case uint8:
		return uint16(rValue), nil
	case uint16:
		return rValue, nil
	case uint32:
		return uint16(rValue), nil
	case uint64:
		return uint16(rValue), nil
	case uintptr:
		return uint16(rValue), nil
	case float32:
		return uint16(rValue), nil
	case float64:
		return uint16(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 16)
		return uint16(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUint32(value interface{}, defaultValue uint32) uint32 {
	var rValue, err = ToUint32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUint32(value interface{}) (uint32, error) {
	switch rValue := value.(type) {
	case int:
		return uint32(rValue), nil
	case int8:
		return uint32(rValue), nil
	case int16:
		return uint32(rValue), nil
	case int32:
		return uint32(rValue), nil
	case int64:
		return uint32(rValue), nil
	case uint:
		return uint32(rValue), nil
	case uint8:
		return uint32(rValue), nil
	case uint16:
		return uint32(rValue), nil
	case uint32:
		return rValue, nil
	case uint64:
		return uint32(rValue), nil
	case uintptr:
		return uint32(rValue), nil
	case float32:
		return uint32(rValue), nil
	case float64:
		return uint32(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 32)
		return uint32(nValue), err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUint64(value interface{}, defaultValue uint64) uint64 {
	var rValue, err = ToUint64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUint64(value interface{}) (uint64, error) {
	switch rValue := value.(type) {
	case int:
		return uint64(rValue), nil
	case int8:
		return uint64(rValue), nil
	case int16:
		return uint64(rValue), nil
	case int32:
		return uint64(rValue), nil
	case int64:
		return uint64(rValue), nil
	case uint:
		return uint64(rValue), nil
	case uint8:
		return uint64(rValue), nil
	case uint16:
		return uint64(rValue), nil
	case uint32:
		return uint64(rValue), nil
	case uint64:
		return rValue, nil
	case uintptr:
		return uint64(rValue), nil
	case float32:
		return uint64(rValue), nil
	case float64:
		return uint64(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return nValue, err
	default:
		return 0, ErrUnsupportedValue
	}
}

func MustUintptr(value interface{}, defaultValue uintptr) uintptr {
	var rValue, err = ToUintptr(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToUintptr(value interface{}) (uintptr, error) {
	switch rValue := value.(type) {
	case int:
		return uintptr(rValue), nil
	case int8:
		return uintptr(rValue), nil
	case int16:
		return uintptr(rValue), nil
	case int32:
		return uintptr(rValue), nil
	case int64:
		return uintptr(rValue), nil
	case uint:
		return uintptr(rValue), nil
	case uint8:
		return uintptr(rValue), nil
	case uint16:
		return uintptr(rValue), nil
	case uint32:
		return uintptr(rValue), nil
	case uint64:
		return uintptr(rValue), nil
	case uintptr:
		return rValue, nil
	case float32:
		return uintptr(rValue), nil
	case float64:
		return uintptr(rValue), nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseUint(trimDecimal(rValue), 10, 64)
		return uintptr(nValue), err
	default:
		return 0, nil
	}
}

func MustString(value interface{}, defaultValue string) string {
	var rValue, err = ToString(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func ToString(value interface{}) (string, error) {
	switch rValue := value.(type) {
	case int:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int8:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int16:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int32:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int64:
		return strconv.FormatInt(rValue, 10), nil
	case uint:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint64:
		return strconv.FormatUint(rValue, 10), nil
	case uintptr:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case float32:
		return strconv.FormatFloat(float64(rValue), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(rValue, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(rValue), nil
	case string:
		return rValue, nil
	case []byte:
		return string(rValue), nil
	case []rune:
		return string(rValue), nil
	default:
		return "", ErrUnsupportedValue
	}
}

func trimDecimal(s string) string {
	if idx := strings.IndexByte(s, '.'); idx != -1 {
		return s[:idx]
	}
	return s
}
