package main

import Server "github.com/justheimsk/vonchat/server/internal/server"

func main() {
  server := Server.NewServer()
  server.Init()

  return
}
