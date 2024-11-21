package types

type UserType struct {
	UserName string `json:"userName"`
	PhoneNo int `json:"phoneNo"`
	UserClass UserClassType `json:"class"`
}

type UserClassType struct {
	Class1 string `json:"class1"`
	Class2 string `json:"class2"`
	UserClassMoreClass UserClassMoreClassType `json:"moreClass"`
}

type UserClassMoreClassType struct {
	Class3 string `json:"class3"`
}

