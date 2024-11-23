package types

type User1Type struct {
	UserName string `json:"userName"`
}
func (q *User1Type) Validate() bool {
	a := true
	return a
}
