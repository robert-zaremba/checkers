package checkers

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
}

var _ = Suite(&S{})

func (s *S) TestToleranceEquality(c *C) {
	c.Check(1.0, EqualsWithTolerance, 1.25, 0.5)
	c.Check(1.0, Not(EqualsWithTolerance), 1.25, 0.05)
}

func (s *S) TestBounds(c *C) {
	c.Check(1.0, Between, 0.0, 1.5)
	c.Check(1.0, Not(Between), 2.0, 2.5)
}

type x struct {
	Val string
}

type y struct {
	Val int
}

func (s *S) TestContains(c *C) {
	a := []int{2, 3, 4}
	c.Check(a, Contains, a[0])
	c.Check(a, Contains, a[1])
	c.Check(a, Contains, a[2])
	c.Check(a, Contains, 2)
	c.Check(a, Contains, 3)
	c.Check(a, Contains, 4)
	c.Check(a, Not(Contains), 5)
	c.Check(a, Not(Contains), a)
	c.Check(a, Not(Contains), "x")

	b := []x{x{"1"}, x{"2"}}
	c.Check(b, Contains, b[0])
	c.Check(b, Contains, b[1])
	c.Check(b, Contains, x{"1"})
	c.Check(b, Contains, x{"2"})
	c.Check(b, Not(Contains), x{"3"})
	c.Check(b, Not(Contains), y{0})

	c.Check("1234", Contains, "23")
	c.Check("1234", Contains, "4")
	c.Check("1234", Contains, "")
	c.Check("1234", Not(Contains), "0")
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
