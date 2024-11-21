package types

type User1 struct {
	UserName string `json:"userName"`
	Id int `json:"id"`
}

func (u *User1) Validate() bool {
	return gt(len(u.UserName), 5) && lt(len(u.UserName), 50) && gt(u.Id, 1) && lt(u.Id, 100)
}

