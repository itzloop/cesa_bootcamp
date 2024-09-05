package professor

import "context"

func Create(ctx context.Context, professor *Professor) (int64, error) {
	var id int64

	return id, nil
}

func GetByID(ctx context.Context, id int64) (*Professor, error) {
	var professor *Professor

	return professor, nil
}

func Update(ctx context.Context, professor *Professor) error {
	return nil
}

func Delete(ctx context.Context, id int64) error {
	return nil
}
