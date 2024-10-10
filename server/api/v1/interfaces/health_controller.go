package interfaces

import "net/http"

type HealthController interface {
	CheckHealth(w http.ResponseWriter, r *http.Request)
}
