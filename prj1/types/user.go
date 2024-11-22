package types

type UserType struct {
	UserName string `json:"userName"`
	PhoneNo int `json:"phoneNo"`
	Class struct {
		Class_1 []string `json:"class 1"`
		Class2 string `json:"class2"`
	} `json:"class"`
}