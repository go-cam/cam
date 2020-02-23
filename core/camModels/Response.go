package camModels

type ResponseModel struct {
	Code    int                    `json:"c"` // status code
	Message string                 `json:"m"` // status message
	Values  map[string]interface{} `json:"v"` // data transferred
}
