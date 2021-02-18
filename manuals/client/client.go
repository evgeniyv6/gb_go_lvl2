package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"

var (
	ErrWrongBodyFmt = errors.New("bad body format")
	ErrSendReq = errors.New("cannot send request")
	UnknownErr = errors.New("unknown error")
	isPanic bool
)


type HttpStatusErr struct {
	status int
}

func NewHttpStatusErr(status int) error {
	return &HttpStatusErr{status}
}

func (e *HttpStatusErr) Error() string {
	return fmt.Sprintf("status code is %d", e.status)
}

func (e *HttpStatusErr) Status() int {
	return e.status
}


func PostJson(client *http.Client, url,body string) (err error) {
	var ijson map[string]interface{}

	defer func() {
		if v:=recover(); v != nil {
			err = fmt.Errorf("%w: %s", UnknownErr, v)
		}
	}()

	fmt.Println("Unmarshal!!!!", ijson)
	if isPanic {
		panic("Wrong json decode")
	}

	if err != nil {
		return fmt.Errorf("%w: %s", ErrWrongBodyFmt, err.Error())
	}

	res, err := client.Post(url, jsonContentType, bytes.NewReader([]byte(body)))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrSendReq, err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return NewHttpStatusErr(res.StatusCode)
	}
	return nil
}
