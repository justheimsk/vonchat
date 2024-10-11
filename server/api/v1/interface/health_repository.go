package interfaces

import "time"

type HealthRepository interface {
	GetPing() (time.Duration, error)
}
