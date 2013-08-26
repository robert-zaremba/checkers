package checkers

import (
	"fmt"
	. "launchpad.net/gocheck"
	"math"
	"reflect"
	"strings"
)

// -----------------------------------------------------------------------
// Contains checker.
type containsChecker struct {
	*CheckerInfo
}

func (c *containsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	container := params[0]
	value := params[1]
	vtype := reflect.TypeOf(value)
	vv := reflect.ValueOf(value)
	cv := reflect.ValueOf(container)

	switch cv.Kind() {
	case reflect.Slice, reflect.Array:
		if cv.Type().Elem() != vtype {
			return false, ""
		}
		for i := 0; i < cv.Len(); i++ {
			if reflect.DeepEqual(cv.Index(i).Interface(), value) {
				return true, ""
			}
		}
		return false, ""
	case reflect.String:
		if vtype.Kind() != reflect.String {
			return false, fmt.Sprint("value should have type: ", vtype)
		}
		return strings.Contains(cv.String(), vv.String()), ""
	}
	return false, fmt.Sprint("Unsupported argument types: ", cv.Kind(), vtype)
}

var Contains Checker = &containsChecker{
	&CheckerInfo{Name: "Contains", Params: []string{"Container", "Value expected to contain"}}}

// -----------------------------------------------------------------------
// EqualsWithTolerance checker.
func equalWithTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func withinBound(value, lower, upper float64) bool {
	return value >= lower && value <= upper
}

type equalsWithToleranceChecker struct {
	*CheckerInfo
}

func (c *equalsWithToleranceChecker) Check(params []interface{}, names []string) (result bool, error string) {
	var (
		ok                            bool
		obtained, expected, tolerance float64
		i                             int64
	)
	obtained, ok = params[0].(float64)
	if !ok {
		i, ok = params[0].(int64)
		if !ok {
			return false, "Obtained value is not a float64 or int64"
		}
		obtained = float64(i)
	}
	expected, ok = params[1].(float64)
	if !ok {
		i, ok = params[1].(int64)
		if !ok {
			return false, "Expected value is not a float64 or int64"
		}
		expected = float64(i)
	}
	tolerance, ok = params[2].(float64)
	if !ok {
		i, ok = params[2].(int64)
		if !ok {
			return false, "Tolerance value is not a float64 or int64"
		}
		tolerance = float64(i)
	}

	return equalWithTolerance(obtained, expected, tolerance), ""
}

var EqualsWithTolerance Checker = &equalsWithToleranceChecker{&CheckerInfo{Name: "EqualsWithTolerance", Params: []string{"obtained", "expected", "tolerance"}}}

// -----------------------------------------------------------------------
// Between checker.
type betweenChecker struct {
	*CheckerInfo
}

func (c *betweenChecker) Check(params []interface{}, names []string) (result bool, error string) {
	var (
		ok                     bool
		obtained, lower, upper float64
	)
	obtained, ok = params[0].(float64)
	if !ok {
		return false, "Obtained value is not a float64"
	}
	lower, ok = params[1].(float64)
	if !ok {
		return false, "Lower value is not a float64"
	}
	upper, ok = params[2].(float64)
	if !ok {
		return false, "Upper value is not a float64"
	}

	return withinBound(obtained, lower, upper), ""
}

var Between Checker = &betweenChecker{&CheckerInfo{Name: "Between", Params: []string{"obtained", "lower", "upper"}}}

// -----------------------------------------------------------------------
// IsTrue checker.
type isTrueChecker struct {
	*CheckerInfo
}

// For example:
//
// c.Assert(v, IsTrue)
var IsTrue Checker = &isTrueChecker{
	&CheckerInfo{Name: "IsTrue", Params: []string{"value"}},
}

func (checker *isTrueChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0] == true, ""
}

// -----------------------------------------------------------------------
// IsFalse checker.
type isFalseChecker struct {
	*CheckerInfo
}

// For example:
//
// c.Assert(v, IsFalse)
var IsFalse Checker = &isFalseChecker{
	&CheckerInfo{Name: "IsFalse", Params: []string{"value"}},
}

func (checker *isFalseChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0] == false, ""
}

// -----------------------------------------------------------------------
// IsEmpty checker.
type isEmptyChecker struct {
	*CheckerInfo
}

// Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or a slice with len == 0.
// For example:
//
// c.Assert(v, IsEmpty)
var IsEmpty Checker = &isEmptyChecker{
	&CheckerInfo{Name: "IsEmpty", Params: []string{"value"}},
}

func (checker *isEmptyChecker) Check(params []interface{}, names []string) (result bool, error string) {
	result = true
	value := params[0]
	if value == nil {
		return
	} else if value == "" {
		return
	} else if value == 0 {
		return
	} else if value == false {
		return
	}

	objValue := reflect.ValueOf(value)
	switch objValue.Kind() {
	case reflect.Slice, reflect.Map:
		return objValue.Len() == 0, ""
	}

	return false, ""
}
