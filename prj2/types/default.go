package types

type RouteType struct {
	TypeName string `json:"typeName"`
	ReqBody interface{} `json:"reqBody"`
	ResBody interface{} `json:"resBody"`
} 

