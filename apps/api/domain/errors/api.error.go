package api_errors

type apiError struct {
	message string
	code    int
}

type ApiErrorPort interface {
	error
	Code() int
}

func New(message string, code int) apiError {
	return apiError{message: message, code: code}
}

func (e apiError) Error() string {
	return e.message
}

func (e apiError) Code() int {
	return e.code
}
