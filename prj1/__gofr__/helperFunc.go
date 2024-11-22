package __gofr__

import (
	"fmt"
	"reflect"
	"strconv"
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

func recurValidate(m *map[string]interface{}, s *string, v *string, level int, parent string) {
	*s += "struct {\n"
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			parts := strings.Fields(val)
			fieldType := parts[0]
			if !isPrimitiveType(fieldType) {
				*v += "a = a && " + "p." + capWord(toUnderscore(key)) + ".Validate()"
				*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " " + capWord(toUnderscore(fieldType)) + "Type `json:\"" + key + "\"`\n"
			} else {
				if len(parts) > 1 {
					*v += "a = a "
				}
				for i := 1; i < len(parts); i++ {
					*v += "&& "
					part := parts[i]
					handleValidatePart(v, part, parent+"."+capWord(toUnderscore(key)))
				}
				*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " " + fieldType + " `json:\"" + key + "\"`\n"
			}
		case map[string]interface{}:
			*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " "
			recurValidate(&val, s, v, level+1, parent+"."+capWord(toUnderscore(key)))
			*s = *s + " `json:\"" + key + "\"`\n"
		case []interface{}:
			if len(val) == 1 {
				switch val1 := val[0].(type) {
				case string:
					*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []" + val1 + " `json:\"" + key + "\"`\n"
				case map[string]interface{}:
					*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []"
					recurValidate(&val1, s, v, level+1, parent+"."+capWord(toUnderscore(key)))
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

func isPrimitiveType(s string) bool {
	return s == "int" || s == "string" || s == "bool" || s == "float64" || s == "float32"
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

func handleValidatePart(v *string, part string, vari string) {
	skipNo := 0
	if strings.HasPrefix(part, "gt=") {
		skipNo = 3
		*v += "gt(" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lt=") {
		skipNo = 3
		*v += "lt(" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "eq=") {
		skipNo = 3
		*v += "eq(" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lengt=") {
		skipNo = 6
		*v += "lengt(len(" + vari + "), " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lenlt=") {
		skipNo = 6
		*v += "lenlt(len(" + vari + "), " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "len=") {
		skipNo = 4
		*v += "len(len(" + vari + "), " + part[skipNo:] + ")"
	}
}
