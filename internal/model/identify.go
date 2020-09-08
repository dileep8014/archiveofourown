package model

import "time"

type Identify struct {
	ID        int64
	Email     string
	Path      string
	CreatedBy string
	CreatedAt time.Time
}

func (ident Identify) TableName() string {
	return "identify"
}
