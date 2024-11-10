package auth

import "net/http"

type Controller interface {
  Register(w http.ResponseWriter, r *http.Request)
  Login(w http.ResponseWriter, r *http.Request)
}
