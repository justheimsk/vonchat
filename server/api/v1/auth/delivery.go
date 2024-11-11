package auth

import "net/http"

type Controller interface {
  Register(http.ResponseWriter, *http.Request)
  Login(http.ResponseWriter, *http.Request)
}
