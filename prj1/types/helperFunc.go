package types

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
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

func gt(a interface{}, b interface{}) bool {
	return convertToFloat64(a) >= convertToFloat64(b)
}

func lt(a interface{}, b interface{}) bool {
	return convertToFloat64(a) <= convertToFloat64(b)
}

func eq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) == convertToFloat64(b)
}

func neq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) != convertToFloat64(b)
}

func lengt(a string, b interface{}) bool {
	return convertToFloat64(len(a)) >= convertToFloat64(b)
}

func lenlt(a string, b interface{}) bool {
	return convertToFloat64(len(a)) <= convertToFloat64(b)
}

func leneq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) == convertToFloat64(b)
}

func lenneq(a interface{}, b interface{}) bool {
	return convertToFloat64(a) != convertToFloat64(b)
}

func req(a interface{}, b bool) bool {
	return a != nil || !b
}

func cateemail(a string) bool {
	re := regexp.MustCompile(`^(?i)[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(a)
}

func cateurl(a string) bool {
	re := `^https?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\\(\\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+$`
	r := regexp.MustCompile(re)
	return r.MatchString(a)
}

func ststr(a string, b string) bool {
	return strings.HasPrefix(a, b)
}
func endstr(a string, b string) bool {
	return strings.HasSuffix(a, b)
}
func constr(a string, b string) bool {
	return strings.Contains(a, b)
}
