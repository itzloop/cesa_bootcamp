package student

import (
	"context"
	"errors"
	"go-university/internal/db"
)

func Create(ctx context.Context, student *Student) (int64, error) {
	if student.ID > 0 {
		return 0, errors.New("student cannot have id in create")
	}

	gormDB := db.GetDB()

	err := gormDB.Create(student).Error
	if err != nil {
		return 0, err
	}

	return student.ID, nil
}

func GetByID(ctx context.Context, id int64) (*Student, error) {
	if id < 1 {
		return nil, errors.New("student id is not valid")
	}

	gormDB := db.GetDB()

	var student *Student

	err := gormDB.Table("student").
		Where("id = ?", id).
		Take(&student).Error
	if err != nil {
		return nil, err
	}

	return student, nil
}

func Update(ctx context.Context, student *Student) error {
	if student.ID < 1 {
		return errors.New("student id is not valid")
	}

	gormDB := db.GetDB()

	err := gormDB.Updates(student).Error
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return errors.New("student id is not valid")
	}

	gormDB := db.GetDB()

	err := gormDB.Where("id = ?", id).
		Delete(&Student{}).Error
	if err != nil {
		return err
	}

	return nil
}
