package camModels

type MessageModel struct {
	Id    int64  `json:"i"` // process id.
	Route string `json:"r"` // route name
	Data  string `json:"d"` // data
}

// Deprecated:
func (model *MessageModel) Validate() bool {
	return model.Route != ""
}
