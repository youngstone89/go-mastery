package gocheck

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

// type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

type MySuite struct {
	dir string
}

func (s *MySuite) SetUpTest(c *C) {
	s.dir = c.MkDir()
	// Use s.dir to prepare some data.
}

func (s *MySuite) TestWithDir(c *C) {
	// Use the data in s.dir in the test.
}
