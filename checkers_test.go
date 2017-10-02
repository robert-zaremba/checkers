package checkers

import (
	"errors"
	"os"

	. "gopkg.in/check.v1"
)

type S struct{}

type x struct {
	Val string
}

type y struct {
	Val int
}

func (s *S) TestIsTrue(c *C) {
	c.Check(true, IsTrue)
	c.Check(false, Not(IsTrue))
	c.Check(1, Not(IsTrue))
	c.Check(nil, Not(IsTrue))
}

func (s *S) TestIsFalse(c *C) {
	c.Check(false, IsFalse)
	c.Check(true, Not(IsFalse))
	c.Check(1, Not(IsFalse))
	c.Check(nil, Not(IsFalse))
}

func (s *S) TestIsEmpty(c *C) {
	c.Check(nil, IsEmpty)
	c.Check(false, IsEmpty)
	c.Check(0, IsEmpty)
	c.Check("", IsEmpty)
	c.Check([]string{}, IsEmpty)
	c.Check(map[int]string{}, IsEmpty)
	c.Check(true, Not(IsEmpty))
	c.Check(1, Not(IsEmpty))
	c.Check("abc", Not(IsEmpty))
	c.Check([]string{"abc", "def"}, Not(IsEmpty))
}

func is42(i int) bool {
	return i == 42
}

var satisfiesTests = []struct {
	f      interface{}
	arg    interface{}
	result bool
	msg    string
}{{
	f:      is42,
	arg:    42,
	result: true,
}, {
	f:      is42,
	arg:    41,
	result: false,
}, {
	f:      is42,
	arg:    "",
	result: false,
	msg:    "wrong argument type string for func(int) bool",
}, {
	f:      os.IsNotExist,
	arg:    errors.New("foo"),
	result: false,
}, {
	f:      os.IsNotExist,
	arg:    os.ErrNotExist,
	result: true,
}, {
	f:      os.IsNotExist,
	arg:    nil,
	result: false,
}, {
	f:      func(chan int) bool { return true },
	arg:    nil,
	result: true,
}, {
	f:      func(func()) bool { return true },
	arg:    nil,
	result: true,
}, {
	f:      func(interface{}) bool { return true },
	arg:    nil,
	result: true,
}, {
	f:      func(map[string]bool) bool { return true },
	arg:    nil,
	result: true,
}, {
	f:      func(*int) bool { return true },
	arg:    nil,
	result: true,
}, {
	f:      func([]string) bool { return true },
	arg:    nil,
	result: true,
}}

func (s *S) TestSatisfies(c *C) {
	for i, test := range satisfiesTests {
		c.Logf("test %d. %T %T", i, test.f, test.arg)
		result, msg := Satisfies.Check([]interface{}{test.arg, test.f}, nil)
		c.Check(result, Equals, test.result)
		c.Check(msg, Equals, test.msg)
	}
}

func (s *S) TestHasPrefix(c *C) {
	c.Assert("foo bar", HasPrefix, "foo")
	c.Assert("foo bar", Not(HasPrefix), "omg")
}

func (s *S) TestHasSuffix(c *C) {
	c.Assert("foo bar", HasSuffix, "bar")
	c.Assert("foo bar", Not(HasSuffix), "omg")
}
