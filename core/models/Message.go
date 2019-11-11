package models

type MessageModel struct {
	BaseModel
	Id    int64  `json:"i"` // 协议号
	Route string `json:"r"` // 路由
	Data  string `json:"d"` // 传输的数据
}

// 验证数据是否合法
func (model *MessageModel) Validate() bool {
	return model.Route != ""
}
