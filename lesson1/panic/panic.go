package panic

import (
	"errors"
	"fmt"
)

type ExtPanic struct {
	txt string
	err error
}

func ErrPrinter() error {
	er := ExtPanic{
		txt: "this is text",
		err: errors.New("new err text"),
	}
	return fmt.Errorf("my custom test %s %w", er.txt, er.err)
}
