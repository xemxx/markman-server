package e

// ErrorData .
type ErrorData struct {
	Code int
	Msg  string
	Err  error
}

func (e *ErrorData) Error() string {
	return e.Msg
}

// Unwrap .
func (e *ErrorData) Unwrap() error { return e.Err }
