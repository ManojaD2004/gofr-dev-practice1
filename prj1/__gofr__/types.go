package __gofr__

type RouteType struct {
	TypeName string `json:"typeName"`
	ReqBody interface{} `json:"reqBody"`
	ResBody interface{} `json:"resBody"`
} 

type NewType struct {
	TypeName string `json:"typeName"`
	TypeBody map[string]interface{} `json:"typeBody"`
} 

type ReturnNewType struct {
	IsCreated bool `json:"isCreated"`
	Message string `json:"message"`
}

type MetaDataType struct {
	Types map[string]map[string]interface{} `json:"types"`
	Routes map[string]map[string]interface{} `json:"routes"`
}

