package camModels

type ResponseModel struct {
	Code    int                    `json:"c"` // 状态码
	Message string                 `json:"m"` // 消息
	Values  map[string]interface{} `json:"v"` // 传输的值
}
