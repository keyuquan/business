package utils

import (
	"fmt"
	"strconv"
)

func InterfaceToInt(s interface{}) int {
	switch s.(type) {
	case int:
		return s.(int)
	case int64:
		return int(s.(int64))
	case int32:
		return int(s.(int32))
	case string:
		i, _ := strconv.Atoi(s.(string))
		return i
	case float64:
		return int(s.(float64))
	case float32:
		return int(s.(float32))
	}
	return 0
}
func InterfaceToFloat64(s interface{}) float64 {
	switch s.(type) {
	case int:
		return float64(s.(int))
	case int64:
		return float64(s.(int64))
	case int32:
		return float64(s.(int32))
	case string:
		f, _ := strconv.ParseFloat(s.(string), 10)
		return f
	case float64:
		return s.(float64)
	case float32:
		return float64(s.(float32))
	}
	return 0
}
func InterfaceToString(s interface{}) string {
	if s == nil {
		return ""
	}
	switch s.(type) {
	case int:
		return strconv.Itoa(s.(int))
	case string:
		return s.(string)
	}
	return fmt.Sprintf("%v", s)
}
