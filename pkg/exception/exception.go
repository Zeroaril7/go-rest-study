package exception

type BaseError struct {
	Error string
}

func NewNotFoundError(error string) BaseError {
	return BaseError{Error: error}
}
