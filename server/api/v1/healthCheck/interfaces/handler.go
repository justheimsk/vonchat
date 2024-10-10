package healthCheckTypes

import (
	"net/http"
)

type Handler interface {
	Load(r *http.ServeMux)
}
