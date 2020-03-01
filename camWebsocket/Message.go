package camWebsocket

type Message struct {
	Id    int64  `json:"i"` // process id.
	Route string `json:"r"` // route name
	Data  string `json:"d"` // data
}
