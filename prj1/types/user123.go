package types

type User123Type struct {
	Friends []string `json:"friends"`
	Marks []struct {
		Mark float64 `json:"mark"`
		MaxMark float32 `json:"maxMark"`
		Subject string `json:"subject"`
	} `json:"marks"`
	Name string `json:"name"`
	Rollno int `json:"rollno"`
}
func (q *User123Type) Validate() bool {
	a := true
	for i1 := 0; i1 < len(q.Friends); i1++ {
		a = a && lengt(q.Friends[i1], 10) && lenlt(q.Friends[i1], 100)
	}
	for i1 := 0; i1 < len(q.Marks); i1++ {
		a = a && lt(q.Marks[i1].MaxMark, 100)
	}
	a = a && lengt(q.Name, 10) && lenlt(q.Name, 100)
	a = a && gt(q.Rollno, 1) && lt(q.Rollno, 100)
	return a
}
