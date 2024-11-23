package __gofr__

import (
	"fmt"
	"reflect"
	// "strconv"
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
				// Validate Part
				*v += "\ta = a && q" + parent + "." + capWord(toUnderscore(key)) + ".Validate()\n"
				// Type Part
				*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " " + capWord(toUnderscore(fieldType)) + "Type `json:\"" + key + "\"`\n"
			} else {
				// Validate Part
				if len(parts) > 1 {
					*v += "\ta = a"
				}
				for i := 1; i < len(parts); i++ {
					*v += " && "
					part := parts[i]
					handleValidatePart(v, part, parent+"."+capWord(toUnderscore(key)))
				}
				if len(parts) > 1 {
					*v += "\n"
				}
				// Type Part
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
					parts := strings.Fields(val1)
					fieldType := parts[0]
					if !isPrimitiveType(fieldType) {
						// Validate Part
						*v += "\tfor i := 0; i < len(q" + parent + "." + capWord(toUnderscore(key)) + "); i++ {\n"
						*v += "\t\ta = a && " + "q" + parent + "." + capWord(toUnderscore(key)) + "[i].Validate()\n"
						*v += "\t}\n"
						// Type Part
						*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []" + capWord(toUnderscore(fieldType)) + "Type `json:\"" + key + "\"`\n"
					} else {
						// Validate Part
						if len(parts) > 1 {
							*v += "\tfor i := 0; i < len(q" + parent + "." + capWord(toUnderscore(key)) + "); i++ {\n"
							*v += "\t\ta = a"
						}
						for i := 1; i < len(parts); i++ {
							*v += " && "
							part := parts[i]
							handleValidatePart(v, part, parent+"."+capWord(toUnderscore(key))+"[i]")
						}
						if len(parts) > 1 {
							*v += "\n\t}\n"
						}
						// Type Part
						*s = *s + strings.Repeat("\t", level) + capWord(toUnderscore(key)) + " []" + fieldType + " `json:\"" + key + "\"`\n"
					}
				case map[string]interface{}:
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
		*v += "Gt(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lt=") {
		skipNo = 3
		*v += "Lt(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "eq=") {
		skipNo = 3
		*v += "Eq(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "neq=") {
		skipNo = 4
		*v += "Neq(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lengt=") {
		skipNo = 6
		*v += "Lengt(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lenlt=") {
		skipNo = 6
		*v += "Lenlt(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "leneq=") {
		skipNo = 6
		*v += "Leneq(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "lenneq=") {
		skipNo = 7
		*v += "Lenneq(q" + vari + ", " + part[skipNo:] + ")"
	} else if strings.HasPrefix(part, "req=") {
		skipNo = 4
		*v += "Req(q" + vari + ", " + part[skipNo:] + ")"
	}
}
