package main
import(
    "regexp"
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