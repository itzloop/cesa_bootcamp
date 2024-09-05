package student

import "time"

type Student struct {
	ID        int64
	FullName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Student) TableName() string {
	return "student"
}
