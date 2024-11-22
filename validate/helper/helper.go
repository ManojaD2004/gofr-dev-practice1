package main

import (
	"regexp"
	"strings"
)

func gt(a, b int) bool {
	return a > b
}

func lt(a, b int) bool {
	return a < b
}
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^(?i)[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func ststr(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}
func endstr(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}
func constr(str, substr string) bool {
	return strings.Contains(str, substr)
}
func url(str string) bool {
	re := `^https?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\\(\\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+$`
	r := regexp.MustCompile(re)
	return r.MatchString(str)
}
func reg(str string, pattern string) bool {
	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return compiledRegex.MatchString(str)
}
