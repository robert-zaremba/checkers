package checkers

import (
	. "gopkg.in/check.v1"
)

type Numeric struct{}

func (s *Numeric) TestToleranceEquality(c *C) {
	c.Check(1.0, EqualsWithTolerance, 1.25, 0.5)
	c.Check(1.0, Not(EqualsWithTolerance), 1.25, 0.05)
}

func (s *Numeric) TestBounds(c *C) {
	c.Check(1.0, Between, 0.0, 1.5)
	c.Check(1.0, Not(Between), 2.0, 2.5)
}
