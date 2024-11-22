package main
import(
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