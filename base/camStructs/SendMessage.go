package camStructs

type SendMessage struct {
	Id    int64       `json:"i"` // process id.
	Route string      `json:"r"` // route name
	Data  interface{} `json:"d"` // response data
}
