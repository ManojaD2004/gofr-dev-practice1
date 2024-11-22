package __gofr__

import (
	"fmt"
	"reflect"
	"strings"
)

func capWord(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func toUnderscore(input string) string {
	result := strings.ReplaceAll(input, "-", "_")
	result = strings.ReplaceAll(result, " ", "_")
	result = strings.ReplaceAll(result, "/", "_")
	return result
}

func recur(m *map[string]interface{}, s *string, level int) {
	*s += "struct {\n"
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " " + val + " `json:\"" + key + "\"`\n"
		case map[string]interface{}:
			*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " "
			recur(&val, s, level+1)
			*s = *s + " `json:\"" + key + "\"`\n"
		case []interface{}:
			if len(val) == 1 {
				switch val1 := val[0].(type) {
				case string:
					*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []" + val1 + " `json:\"" + key + "\"`\n"
				case map[string]interface{}:
					*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []"
					recur(&val1, s, level+1)
					*s = *s + " `json:\"" + key + "\"`\n"
				default:
					fmt.Println("Error 2!", reflect.TypeOf(val1))
				}
			}
		default:
			fmt.Println("Error 1!", reflect.TypeOf(val))
		}
	}
	*s += strings.Repeat("\t", level-1) + "}"
}

func convertStructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)
	if val.Kind() != reflect.Struct {
		panic("input must be a struct")
	}
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		result[fieldName] = fieldValue
	}
	return result
}
