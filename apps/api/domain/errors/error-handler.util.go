package api_errors

func ErrorHandler[T any](input T, e error) (T, ApiErrorPort) {
	if e != nil {
		message := e.Error()
		var statusCode int
		if apiErr, ok := e.(ApiErrorPort); ok {
			statusCode = apiErr.Code()
		} else {
			statusCode = 500
		}

		if len(message) > 0 {
			return input, New(message, statusCode)
		}
		return input, New("Internal Server Error", statusCode)
	}
	return input, nil
}
