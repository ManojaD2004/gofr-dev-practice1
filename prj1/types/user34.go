package types

type User34Type struct {
	IsHandled string `json:"isHandled"`
}
func (q *User34Type) Validate() bool {
	a := true
	a = a && lengt(q.IsHandled, 50)
	return a
}
