package professor

import "time"

type Professor struct {
	ID        int64
	FullName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Professor) TableName() string {
	return "professor"
}
