package exception

type notFoundError struct {
	Error string
}

func NewNotFoundError(err string) notFoundError {
	return notFoundError{Error: err}
}
