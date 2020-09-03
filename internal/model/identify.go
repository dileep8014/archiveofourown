package model

import "time"

type Identify struct {
	ID        int64
	Email     string
	CreatedBy string
	CreatedAt time.Time
}
