package healthCheck

import "net/http"

type Controller interface {
  CheckHealth(http.ResponseWriter, *http.Request)
}
