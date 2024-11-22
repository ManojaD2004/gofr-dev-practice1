package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
func recur(m *map[string]interface{}, key1 string, s *string, validateLogic *string, parent string) {
	structDef := "type " + capitalizeFirstLetter(key1) + "Type struct {\n"
	validationLogic := ""

	for key, value := range *m {
		switch val := value.(type) {
		case string:
			var fieldType string
			var gt, lt, length int
			var startStr, endStr, containsStr, regexPattern string
			gtProvided, ltProvided, lengthProvided, startStrProvided, endStrProvided, containsStrProvided, regexProvided := false, false, false, false, false, false, false

			parts := strings.Fields(val)
			fieldType = parts[0]

			for _, part := range parts[1:] {
				if strings.HasPrefix(part, "gt=") {
					fmt.Sscanf(part, "gt=%d", &gt)
					gtProvided = true
				} else if strings.HasPrefix(part, "lt=") {
					fmt.Sscanf(part, "lt=%d", &lt)
					ltProvided = true
				} else if strings.HasPrefix(part, "length=") {
					fmt.Sscanf(part, "length=%d", &length)
					lengthProvided = true
				} else if strings.HasPrefix(part, "startswith=") {
					fmt.Sscanf(part, "startswith=%s", &startStr)
					startStrProvided = true
				} else if strings.HasPrefix(part, "startswith=") {
					fmt.Sscanf(part, "startswith=%s", &endStr)
					endStrProvided = true
				} else if strings.HasPrefix(part, "contains=") {
					fmt.Sscanf(part, "contains=%s", &containsStr)
					containsStrProvided = true
				} else if strings.HasPrefix(part, "regex=") {
					fmt.Sscanf(part, "regex=%s", &regexPattern)
					regexProvided = true
				}
			}
			structDef += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", capitalizeFirstLetter(key), fieldType, key)
			fieldPath := parent + capitalizeFirstLetter(key)
			if fieldType == "string" {
				if lengthProvided {
					validationLogic += fmt.Sprintf("len(%s) == %d && ", fieldPath, length)
				} else {
					if gtProvided {
						validationLogic += fmt.Sprintf("gt(len(%s), %d) && ", fieldPath, gt)
					}
					if ltProvided {
						validationLogic += fmt.Sprintf("lt(len(%s), %d) && ", fieldPath, lt)
					}

				}
				if startStrProvided {
					validationLogic += fmt.Sprintf("ststr(%s, %s) && ", fieldPath, startStr)
				}
				if endStrProvided {
					validationLogic += fmt.Sprintf("endstr(%s, %s) && ", fieldPath, endStr)
				}
				if containsStrProvided {
					validationLogic += fmt.Sprintf("constr(%s, %s) && ", fieldPath, containsStr)
				}
				if regexProvided {
					validationLogic += fmt.Sprintf("reg(%s, \"%s\") && ", fieldPath, regexPattern)
				}
				if key == "email" {
					validationLogic += fmt.Sprintf("em(%s) && ", fieldPath)
				}
				if key == "url" {
					validationLogic += fmt.Sprintf("url(%s) && ", fieldPath)
				}

			} else {
				if gtProvided {
					validationLogic += fmt.Sprintf("gt(%s, %d) && ", fieldPath, gt)
				}
				if ltProvided {
					validationLogic += fmt.Sprintf("lt(%s, %d) && ", fieldPath, lt)
				}
			}

		case map[string]interface{}:
			nestedType := capitalizeFirstLetter(key1 + capitalizeFirstLetter(key))
			structDef += fmt.Sprintf("\t%s %sType `json:\"%s\"`\n", capitalizeFirstLetter(key), nestedType, key)
			recur(&val, nestedType, s, validateLogic, parent+capitalizeFirstLetter(key)+".")
		default:
			fmt.Println("Error!")
		}
	}
	structDef += "}\n\n"
	*s = structDef + *s
	*validateLogic += validationLogic
}
func main() {
	Schema := `{
		"empName": "string gt=5 lt=50",
		"id": "int gt=1 lt=100",
		"email": "string gt=10 lt=100",
		"address": {
		  "street": "string gt=10 lt=100",
		  "city": "string gt=3 lt=50",
		  "zip": "int lt=9999",
		  "country": {
			"name": "string gt=3 lt=50",
			"code": "string gt=2 lt=10"
		  }
		},
		"preferences": {
		  "theme": "string gt=3 lt=20",
		  "notificationsEnabled": "bool",
		  "language": "string gt=2 lt=10"
		}
	  }`
	var schema map[string]interface{}
	json.Unmarshal([]byte(Schema), &schema)
	code := ""
	validateLogic := ""
	recur(&schema, "User", &code, &validateLogic, "u.")
	validateFunction := "func Validate(u *UserType) bool {\n\treturn " + validateLogic[:len(validateLogic)-4] + "\n}\n\n"
	code = validateFunction + code
	fmt.Println(code)
}
