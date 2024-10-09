package server

import "fmt"

type Server struct {}

func NewServer() (*Server) {
  return &Server{}
}

func (self *Server) Init() {
  fmt.Println("Init server");
}
