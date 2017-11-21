package checkers

import (
	"fmt"
	"time"

	gc "gopkg.in/check.v1"
)

// TimeBetween returns a time between checker
func TimeBetween(start, end time.Time) gc.Checker {
	if end.Before(start) {
		return &timeBetweenChecker{end, start}
	}
	return &timeBetweenChecker{start, end}
}

type timeBetweenChecker struct {
	start, end time.Time
}

func (checker *timeBetweenChecker) Info() *gc.CheckerInfo {
	info := gc.CheckerInfo{
		Name:   "TimeBetween",
		Params: []string{"obtained"},
	}
	return &info
}

func (checker *timeBetweenChecker) Check(params []interface{}, names []string) (result bool, error string) {
	when, ok := params[0].(time.Time)
	if !ok {
		return false, "obtained value type must be time.Time"
	}
	if when.Before(checker.start) {
		return false, fmt.Sprintf("obtained value %#v type must before start value of %#v", when, checker.start)
	}
	if when.After(checker.end) {
		return false, fmt.Sprintf("obtained value %#v type must after end value of %#v", when, checker.end)
	}
	return true, ""
}

// -----------------------------------------------------------------------
type durationLessThanChecker struct {
	*gc.CheckerInfo
}

// DurationLessThan checker
var DurationLessThan gc.Checker = &durationLessThanChecker{
	&gc.CheckerInfo{Name: "DurationLessThan", Params: []string{"obtained", "expected"}},
}

func (checker *durationLessThanChecker) Check(params []interface{}, names []string) (result bool, error string) {
	obtained, ok := params[0].(time.Duration)
	if !ok {
		return false, "obtained value type must be time.Duration"
	}
	expected, ok := params[1].(time.Duration)
	if !ok {
		return false, "expected value type must be time.Duration"
	}
	return obtained.Nanoseconds() < expected.Nanoseconds(), ""
}

// -----------------------------------------------------------------------
type withinDuration struct {
	*gc.CheckerInfo
}

func (checker *withinDuration) Check(params []interface{}, names []string) (result bool, error string) {
	obtained, expected, errstr := toTime(params[0], params[1])
	if errstr != "" {
		return false, errstr
	}
	maxDiff, ok := params[2].(time.Duration)
	if !ok {
		return false, "max_diff value type must be time.Duration"
	}
	dt := expected.Sub(obtained)
	if dt >= -maxDiff && dt <= maxDiff {
		return true, ""
	}
	return false, ""
}

// WithinDuration checkes if time between obtained and expected is within duration
var WithinDuration gc.Checker = &withinDuration{
	&gc.CheckerInfo{Name: "WithinDuration", Params: []string{"obtained", "expected", "max_diff"}}}

// -----------------------------------------------------------------------
type timeEquals struct {
	*gc.CheckerInfo
}

func (checker *timeEquals) Check(params []interface{}, names []string) (result bool, errstr string) {
	obtained, expected, errstr := toTime(params[0], params[1])
	if errstr != "" {
		return false, errstr
	}
	maxDiff := time.Microsecond
	dt := expected.Sub(obtained)
	if dt >= -maxDiff && dt <= maxDiff {
		return true, ""
	}
	return false, ""
}

// TimeEquals checkes if time between obtained and expected is within duration
var TimeEquals gc.Checker = &timeEquals{
	&gc.CheckerInfo{Name: "TimeEquals", Params: []string{"obtained", "expected", "max_diff"}},
}

// -----------------------------------------------------------------------

func toTime(a, b interface{}) (time.Time, time.Time, string) {
	obtained, ok := a.(time.Time)
	if !ok {
		return obtained, obtained, "obtained value type must be time.Time"
	}
	expected, ok := b.(time.Time)
	if !ok {
		return obtained, expected, "expected value type must be time.Time"
	}
	return obtained, expected, "expected value type must be time.Time"
}
