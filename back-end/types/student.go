package types

type StudentType struct {
	Name string `json:"name"`
	Rollno int `json:"rollno"`
	Marks []struct {
		Subject string `json:"subject"`
		Mark float64 `json:"mark"`
		MaxMark float32 `json:"maxMark"`
	} `json:"marks"`
	Friends []string `json:"friends"`
}
func (q *StudentType) Validate() bool {
	a := true
	a = a && lengt(q.Name, 10) && lenlt(q.Name, 100)
	a = a && gt(q.Rollno, 1) && lt(q.Rollno, 100)
	for i1 := 0; i1 < len(q.Marks); i1++ {
		a = a && lt(q.Marks[i1].MaxMark, 100)
	}
	for i1 := 0; i1 < len(q.Friends); i1++ {
		a = a && lengt(q.Friends[i1], 10) && lenlt(q.Friends[i1], 100)
	}
	return a
}
