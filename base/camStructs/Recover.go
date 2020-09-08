package camStructs

import (
	"errors"
	"github.com/go-cam/cam/base/camStatics"
)

// recoverable panic content
type Recover struct {
	camStatics.RecoverInterface
	error
}

// new recoverable
func NewRecover(message string) *Recover {
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
