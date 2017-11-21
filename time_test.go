package checkers

import (
	"time"

	. "gopkg.in/check.v1"
)

type Time struct{}

func (ts *Time) TestWithinDuration(c *C) {
	t1 := time.Now()
	maxDiff := time.Minute

	c.Check(t1, WithinDuration, t1.Add(time.Second), maxDiff)
	c.Check(t1, WithinDuration, t1.Add(-time.Second), maxDiff)
	c.Check(t1, WithinDuration, t1.Add(time.Minute), maxDiff)
	c.Check(t1, WithinDuration, t1.Add(-time.Minute), maxDiff)

	c.Check(t1, Not(WithinDuration), t1.Add(2*time.Minute), maxDiff)
	c.Check(t1, Not(WithinDuration), t1.Add(-2*time.Minute), maxDiff)
}

func (ts *Time) TestTimeEquals(c *C) {
	t1 := time.Now()

	c.Check(t1, TimeEquals, t1)
	c.Check(t1, TimeEquals, t1.UTC())
	c.Check(t1, TimeEquals, t1.Add(time.Nanosecond*10))
	c.Check(t1, TimeEquals, t1.Add(-time.Nanosecond*120))

	c.Check(t1, Not(TimeEquals), t1.Add(time.Microsecond+1))
}
