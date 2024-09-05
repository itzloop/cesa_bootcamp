package grade

import (
	"time"

	"github.com/shopspring/decimal"
)

type Grade struct {
	ID        int64
	StudentID int64
	CourseID  int64
	Grade     decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Grade) TableName() string {
	return "grade"
}
