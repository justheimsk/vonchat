package healthResourceType

import "time"

type Repository interface {
	GetPing() (time.Duration, error)
}
