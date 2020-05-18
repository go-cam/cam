package camStructs

type Message struct {
	Id    int64                  `json:"i"` // process id.
	Route string                 `json:"r"` // route name
	Data  map[string]interface{} `json:"d"` // data
}
