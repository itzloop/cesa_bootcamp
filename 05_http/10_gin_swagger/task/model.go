package task

type GetAllResponse struct {
	Items []*Task
}

type Task struct {
	ID    int64
	Title string
}

type GetAllRequest struct {
	Page int64
	Size int64
	Term string
}

type DeleteRequest struct {
	ID int64
}

type UpdateRequest struct {
	ID    int64
	Title string
}

type GetResponse struct {
	ID    int64
	Title string
}

type GetRequest struct {
	ID int64
}

type CreateResponse struct {
	ID    int64
	Title string
}

type CreateRequest struct {
	Title string
}
