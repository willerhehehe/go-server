package mysql

type DBError struct {
	msg string
}

func (e *DBError) Error() string {
	return e.msg
}

func NewDBError(msg string) error {
	return &DBError{msg}
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "data not found"
}
