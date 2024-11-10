package domain_repo

import "time"

type HealthRepository interface {
	GetPing() (time.Duration, error)
}
