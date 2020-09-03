package model

import "time"

type Subsection struct {
	ID        int64
	WorkId    int64
	Name      string
	Introduce string
	Seq       int
	WorkNum   int
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
