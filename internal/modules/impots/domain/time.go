package domain

import "time"

type Time interface {
	Now() time.Time
}
