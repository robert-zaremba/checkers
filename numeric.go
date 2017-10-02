package checkers

import (
	"fmt"
	"math"

	gc "gopkg.in/check.v1"
)

func toFloat(val interface{}) (float64, string) {
	switch t := val.(type) {
	case float64:
		return t, ""
	case float32:
		return float64(t), ""
	case int:
		return float64(t), ""
	case int32:
		return float64(t), ""
	case int64:
		return float64(t), ""
	}
	return 0, fmt.Sprintf("Expecting a number, got: %T", val)
}

// -----------------------------------------------------------------------
func equalWithTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func withinBound(value, lower, upper float64) bool {
	return value >= lower && value <= upper
}

type equalsWithToleranceChecker struct {
	*gc.CheckerInfo
}

func (c *equalsWithToleranceChecker) Check(params []interface{}, names []string) (result bool, error string) {
	var obtained, expected, tolerance float64
	var errStr string
	if obtained, errStr = toFloat(params[0]); errStr != "" {
		return false, "Wrong obtained value: " + errStr
	}
	if expected, errStr = toFloat(params[1]); errStr != "" {
		return false, "Wrong expected value: " + errStr
	}
	if tolerance, errStr = toFloat(params[2]); errStr != "" {
		return false, "Wrong tolerance value: " + errStr
	}
	return equalWithTolerance(obtained, expected, tolerance), ""
}

// EqualsWithTolerance Check if two numbers are close enough
var EqualsWithTolerance gc.Checker = &equalsWithToleranceChecker{&gc.CheckerInfo{Name: "EqualsWithTolerance", Params: []string{"obtained", "expected", "tolerance"}}}

// CloseTo is an alias for EqualsWithTolerance
var CloseTo = EqualsWithTolerance

// -----------------------------------------------------------------------
type betweenChecker struct {
	*gc.CheckerInfo
}

// Between checks if a numeric values is between two other values
var Between gc.Checker = &betweenChecker{&gc.CheckerInfo{Name: "Between", Params: []string{"obtained", "lower", "upper"}}}

func (c *betweenChecker) Check(params []interface{}, names []string) (result bool, error string) {
	var obtained, upper, lower float64
	var errStr string
	if obtained, errStr = toFloat(params[0]); errStr != "" {
		return false, "Wrong obtained value: " + errStr
	}
	if lower, errStr = toFloat(params[1]); errStr != "" {
		return false, "Wrong lower value: " + errStr
	}
	if upper, errStr = toFloat(params[2]); errStr != "" {
		return false, "Wrong upper value: " + errStr
	}
	return withinBound(obtained, lower, upper), ""
}
