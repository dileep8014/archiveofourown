package model

import "time"

type Subscription struct {
	ID        int64
	UserId    int64
	ObjType   int
	ObjId     int64
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
