package types

type StudentsType struct {
	AllStudents []StudentType `json:"allStudents"`
}
func (q *StudentsType) Validate() bool {
	a := true
	for i1 := 0; i1 < len(q.AllStudents); i1++ {
		a = a && q.AllStudents[i1].Validate()
	}
	return a
}
