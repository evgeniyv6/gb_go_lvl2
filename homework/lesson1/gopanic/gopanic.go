package gopanic

import (
	"errors"
	"fmt"
	"time"
)


// ErrWithTimestamp desc
type ErrWithTimestamp struct {
	text string
	datetime string
}


// WrapErrWithTimeStamp - try wrap error
type WrapErrWithTimeStamp struct {
	msg string
	err error
}

func (e *ErrWithTimestamp) Error() string {
	return fmt.Sprintf("Error: %s occured at %v", e.text, e.datetime)
}

func (e *WrapErrWithTimeStamp) Error() string {
	return e.msg
}


// New desc
func New(msg string) error {
	return &ErrWithTimestamp{msg, time.Now().Format(time.RFC1123)}
}

func devider(a,b int) {
	fmt.Println(a / b)
}

// GoroutinePanicCatcher description (for linter)
func GoroutinePanicCatcher() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(New("catch panic in goroutine"))
				var err error
				err = errors.New("my error")
				err = fmt.Errorf("wrapped error: %w", err)
				fmt.Println(err)
			}
		}()
		devider(1, 0)
	}()

	time.Sleep(1 * time.Second)
}
