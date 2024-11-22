package types

import (
	"fmt"
	"reflect"
)

func convertToFloat64(a interface{}) float64 {
	var a2 float64
	switch a1 := a.(type) {
	case int:
		a2 = float64(a1)
	case float32:
		a2 = float64(a1)
	case float64:
		a2 = a1
	default:
		a2 = 0.0
		fmt.Println("Wrong type, only int, float32, float64 ", reflect.TypeOf(a1))
	}
	return a2
}

func Gt(a interface{}, b interface{}) bool {
	return convertToFloat64(a) >= convertToFloat64(b)
}

func Lt(a interface{}, b interface{}) bool {
	return convertToFloat64(a) <= convertToFloat64(b)
}

func Eq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) == convertToFloat64(b)
}

func Neq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) != convertToFloat64(b)
}

func Lengt(a string, b interface{}) bool {
	return convertToFloat64(len(a)) >= convertToFloat64(b)
}

func Lenlt(a string, b interface{}) bool {
	return convertToFloat64(len(a)) <= convertToFloat64(b)
}

func Leneq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) == convertToFloat64(b)
}

func Lenneq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) != convertToFloat64(b)
}

func Req(a interface{}, b bool) bool {
	return a != nil || !b
}

