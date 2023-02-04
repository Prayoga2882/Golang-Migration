package exception

type NotFoundError struct {
	Error string
}

// fix_this_shit !!!

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
