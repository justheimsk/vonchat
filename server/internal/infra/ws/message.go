package ws

type MSG_T string
type MSG_OP int
type MSG_DATA any

type WebsocketMessage struct {
	Op   MSG_OP   `json:"op"`
	Data MSG_DATA `json:"d"`
	T    MSG_T    `json:"t"`
}

func NewWebsocketMessage(op MSG_OP, t MSG_T, data MSG_DATA) *WebsocketMessage {
	return &WebsocketMessage{
		Op:   op,
		Data: data,
		T:    t,
	}
}
