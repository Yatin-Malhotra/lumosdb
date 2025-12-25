package storage

import "time"

type Item struct {
	Value     string
	ExpireAt  time.Time
	HasExpiry bool
}
