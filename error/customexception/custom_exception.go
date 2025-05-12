package customexception

import "fmt"

func New(code int, message string) error {
	return &customExceptionError{code, message}
}

type customExceptionError struct {
	code    int
	message string
}

func (b *customExceptionError) Error() string {
	return fmt.Sprintf("%d - %s", b.code, b.message)
}
