package safe

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"runtime"
)

type PanicError struct {
	// The function that was given to `Call()`.
	In func() error
	// The recovered panic value while executing `In()`.
	Err error
}

func NewPanicError(in func() error, err error) *PanicError {
	return &PanicError{
		In:  in,
		Err: errors.WithStack(err),
	}
}

func (e *PanicError) Unwrap() error {
	return e.Err
}

func (e *PanicError) Cause() error {
	return e.Err
}

func (e *PanicError) inName() string {
	return runtime.FuncForPC(reflect.ValueOf(e.In).Pointer()).Name()
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic while executing %s: %#+v", e.inName(), e.Err)
}

func Call(f func() error)(err error)  {
	defer func() {
		r := recover()
		if r == nil {
			// Note that panic(nil) matches this case and cannot be really tested for.
			return
		}

		switch actual := r.(type) {
		case error:
			err = actual
		case string:
			err = errors.New(actual)
		default:
			err = errors.Errorf("%v", r)
		}
		fmt.Println("[!] call.go line:56 [",err,"]")
		err = NewPanicError(f, err)
	}()

	return f()
}