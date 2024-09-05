package student

import "context"

func Create(ctx context.Context, student *Student) (int64, error) {
	var id int64

	return id, nil
}

func GetByID(ctx context.Context, id int64) (*Student, error) {
	var student *Student

	return student, nil
}

func Update(ctx context.Context, student *Student) error {
	return nil
}

func Delete(ctx context.Context, id int64) error {
	return nil
}
