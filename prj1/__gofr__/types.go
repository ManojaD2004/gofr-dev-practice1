package __gofr__

type RouteType struct {
	TypeName string `json:"typeName"`
	ReqBody interface{} `json:"reqBody"`
	ResBody interface{} `json:"resBody"`
} 

type NewType struct {
	TypeName string `json:"typeName"`
	TypeBody interface{} `json:"typeBody"`
} 

type ReturnNewType struct {
	IsCreated bool `json:"isCreated"`
	Message string `json:"message"`
}

