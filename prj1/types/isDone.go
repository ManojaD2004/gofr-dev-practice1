package types

type IsDoneType struct {
	IsHandled bool `json:"isHandled"`
	Message string `json:"message"`
}
func (q *IsDoneType) Validate() bool {
	a := true
	a = a && lengt(q.Message, 10)
	return a
}
