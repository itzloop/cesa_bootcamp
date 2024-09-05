package student

import "time"

type Student struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (Student) TableName() string {
	return "student"
}
