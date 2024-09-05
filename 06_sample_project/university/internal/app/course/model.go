package course

import "time"

type Course struct {
	ID        int64
	Title     string
	Credits   int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Course) TableName() string {
	return "course"
}
