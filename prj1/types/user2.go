package types

type User2Type struct {
	Class struct {
		Class2 string `json:"class2"`
		More_Class []struct {
			School4 []struct {
				Class5 string `json:"class5"`
			} `json:"school4"`
			Class3 string `json:"class3"`
		} `json:"more-Class"`
		Class_1 []int `json:"class 1"`
	} `json:"class"`
	UserName string `json:"userName"`
	PhoneNo int `json:"phoneNo"`
}

func (q *User2Type) Validate() bool {
	return q.PhoneNo == 00
}