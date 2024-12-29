package errors

// AuthError is a special error type used to signal an authorization error
type AuthError struct {
	Err error
}

// Error satisfies error interface for AuthError
func (a AuthError) Error() string {
	return a.Err.Error()
}
