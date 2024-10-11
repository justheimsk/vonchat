package domain

import "time"

type HealthRepository interface {
	GetPing() (time.Duration, error)
}
