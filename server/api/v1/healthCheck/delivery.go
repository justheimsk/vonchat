package healthCheck

import "net/http"

type Controller interface {
  CheckHealth(w http.ResponseWriter, r *http.Request)
}
