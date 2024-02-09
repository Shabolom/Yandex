package customErr

import (
	"fmt"
)

const ERR_POST = "уже имеется в базе"

type Err struct {
	err  error
	code int
}

func (e *Err) Error() string {
	return fmt.Sprintf("err: \v // code err: \v", e.err, e.code)
}

func NewErr(err error, code int) error {
	return &Err{
		err:  err,
		code: code,
	}
}
