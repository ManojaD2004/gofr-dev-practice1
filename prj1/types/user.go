package types

type UserType struct {
	UserName string `json:"userName"`
	PhoneNo  int    `json:"phoneNo"`
	Class    struct {
		Class1    []string `json:"class1"`
		Class2    string   `json:"class2"`
		MoreClass []struct {
			Class3  string `json:"class3"`
			School4 []struct {
				Class5 string `json:"class5"`
			} `json:"school4"`
		} `json:"moreClass"`
	} `json:"class"`
}
