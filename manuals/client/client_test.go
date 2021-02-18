package client

import (
	"errors"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func Test (t *testing.T) { TestingT(t) }

type TestSuite struct {}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestWrongJson(c *C) {
	var err error
	err = PostJson(http.DefaultClient, "http://google.com", `{{"data": 1}`)
	c.Assert(err, NotNil)
	c.Assert(errors.Is(err,ErrWrongBodyFmt), Equals, true)
}


func (s *TestSuite) TestWrongUrl (c *C) {
	var err error
	err = PostJson(http.DefaultClient, "http://icorrect.url", `{}`)
	c.Assert(err, NotNil)
	c.Assert(errors.Is(err,ErrSendReq), Equals, true)
}

func (s *TestSuite) TestWrongStatus (c *C) {
	var err error
	err = PostJson(http.DefaultClient, "http://httpstat.us/500", `{"data": {}}`)
	c.Assert(err, NotNil)
	e, ok := err.(*HttpStatusErr)
	c.Assert(ok, Equals, true)
	c.Assert(e.Status(), Equals, http.StatusInternalServerError)
}

func (s *TestSuite) TestOk (c *C) {
	err:= PostJson(http.DefaultClient, "http://httpstat.us/200", `{}`)
	c.Assert(err, IsNil)

}

func (s *TestSuite) TestUnknownErr (c *C) {
	isPanic = true
	var err error
	err = PostJson(http.DefaultClient,"http://httpstat.us/200", `{}`)
	c.Assert(err, NotNil)
	c.Assert(errors.Is(err, UnknownErr), Equals, true)

	isPanic = false
}







