package camBase

import "errors"

// recoverable panic content
type Recover struct {
	RecoverInterface
	error
}

// new recoverable
func NewRecoverable(message string) *Recover {
	r := new(Recover)
	r.error = errors.New(message)
	return r
}

// get error string
func (r *Recover) Error() string {
	return r.error.Error()
}

// get error
func (r *Recover) GetError() error {
	return r.error
}
