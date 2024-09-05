package course

import "time"

type Course struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Credits   int32  `json:"credits"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CourseProfessor struct {
	ID          int64
	ProfessorID int64
	CourseID    int64
}

func (Course) TableName() string {
	return "course"
}

func (CourseProfessor) TableName() string {
	return "course_professor"
}
