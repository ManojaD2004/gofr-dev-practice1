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

func smallFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func toUnderscore(input string) string {
	result := strings.ReplaceAll(input, "-", "_")
	result = strings.ReplaceAll(result, " ", "_")
	return result
}

func recur(m *map[string]interface{}, s *string, level int) {
	*s += "struct {\n"
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			*s = *s + strings.Repeat("\t", level) + capWord(key) + " " + val + " `json:\"" + key + "\"`\n"
			// fmt.Println(key1, key, val, s)
		case map[string]interface{}:
			// fmt.Println(key1, key, val)
			newType := capWord(key)
			*s = *s + strings.Repeat("\t", level) + newType + " "
			recur(&val, s, level+1)
			*s = *s + " `json:\"" + key + "\"`\n"
		case []interface{}:
			if len(val) == 1 {
				switch val1 := val[0].(type) {
				case string:
					*s = *s + strings.Repeat("\t", level) + capWord(key) + " []" + val1 + " `json:\"" + key + "\"`\n"
					// fmt.Println("Array of string")
				case map[string]interface{}:
					// fmt.Println("Array of object")
					newType := capWord(key)
					*s = *s + strings.Repeat("\t", level) + newType + " []"
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
