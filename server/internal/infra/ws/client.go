package ws

import (
	"github.com/gorilla/websocket"
	"github.com/justheimsk/vonchat/server/internal/application/dto"
)

type Client struct {
	RandomID      string
	Socket        *websocket.Conn
	Authenticated bool
	user          *dto.UserDTO
}

func NewClient(id string, socket *websocket.Conn) *Client {
	return &Client{
		RandomID:      id,
		Socket:        socket,
		Authenticated: false,
	}
}

func (self *Client) Authenticate(user *dto.UserDTO) {
	self.user = user
	self.Authenticated = true
}
