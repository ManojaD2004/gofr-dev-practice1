package types

type User2Type struct {
	UserName string `json:"userName"`
	PhoneNo int `json:"phoneNo"`
	Class struct {
		Class_1 []string `json:"class 1"`
		Class2 string `json:"class2"`
		More_Class []struct {
			Class3 string `json:"class3"`
			School4 []struct {
				Class5 string `json:"class5"`
			} `json:"school4"`
		} `json:"more-Class"`
	} `json:"class"`
}