package __gofr__

type CreateRouteType struct {
	RouteName   string `json:"routeName"`
	Method string `json"method"`
	ReqBodyType string `json:"reqBodyType"`
	ResBodyType string `json:"resBodyType"`
}

type DeleteRouteType struct {
	RouteName   string `json:"routeName"`
	Method string `json"method"`
}

type NewType struct {
	TypeName string                 `json:"typeName"`
	TypeBody map[string]interface{} `json:"typeBody"`
}

type DeleteType struct {
	TypeName string `json:"typeName"`
}

type GetType struct {
	TypeName string `json:"typeName"`
}

type ReturnType struct {
	IsDone  bool        `json:"isDone"`
	Message interface{} `json:"message"`
}

type MetaDataType struct {
	Types  map[string]map[string]interface{} `json:"types"`
	Routes map[string]map[string]interface{} `json:"routes"`
}
