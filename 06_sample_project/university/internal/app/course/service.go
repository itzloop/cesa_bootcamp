package course

import (
	"context"
	"errors"
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
		Take(&course).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func Update(ctx context.Context, course *Course) error {
	if course.ID < 1 {
		return errors.New("course is not valid")
	}

	gormDB := db.GetDB()

	err := gormDB.Updates(&course).Error
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return errors.New("course id is not valid")
	}

	gormDB := db.GetDB()

	err := gormDB.Where("id = ?", id).
		Delete(&Course{}).Error
	if err != nil {
		return err
	}

	return nil
}
