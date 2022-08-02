package errs

import "errors"

var ErrNotFound = errors.New("not found")

var _ error = new(NotFoundError)

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + ": not found"
}

func NewNotFoundError() error {
	return &NotFoundError{}
}

type QueryError struct {
	Query string
	Err   error
}

// Error implements error
func (q QueryError) Error() string {
	return q.Query + q.Err.Error()
}

func NewQueryError() error {
	return QueryError{Query: "query", Err: ErrNotFound}
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

type PermissionError struct {
}

func (pe *PermissionError) Error() string {
	return "permission error"
}
