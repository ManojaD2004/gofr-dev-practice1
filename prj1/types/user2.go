package types

type User2Type struct {
	AllStudent []User1Type `json:"allStudent"`
}
func (q *User2Type) Validate() bool {
	a := true
	for i1 := 0; i1 < len(q.AllStudent); i1++ {
		a = a && q.AllStudent[i1].Validate()
	}
	return a
}
