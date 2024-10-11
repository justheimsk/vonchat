package interfaces

import (
	"net/http"
	"time"
)

type HealthController interface {
	CheckHealth(w http.ResponseWriter, r *http.Request)
}

type HealthRepository interface {
	GetPing() (time.Duration, error)
}
