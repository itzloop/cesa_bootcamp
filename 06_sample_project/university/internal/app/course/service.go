package course

import (
	"context"
	"go-university/internal/db"
)

func Create(ctx context.Context, course *Course) (int64, error) {
	gormDB := db.GetDB()

	err := gormDB.Create(course).Error
	if err != nil {
		return 0, err
	}

	return course.ID, nil
}

func GetByID(ctx context.Context, id int64) (*Course, error) {
	gormDB := db.GetDB()

	var course *Course
	err := gormDB.Model(&Course{}).
		Where("id = ?", id).
		Scan(&course).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func Update(ctx context.Context, course *Course) error {
	return nil
}

func Delete(ctx context.Context, id int64) error {
	return nil
}
