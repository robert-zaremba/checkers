package checkers

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	. "gopkg.in/check.v1"
	"math"
	"reflect"
	"strings"
)

type comment struct {
	args   []interface{}
	isSpew bool
}

// Comment is like Commentf but without formatting string
func Comment(args ...interface{}) CommentInterface {
	return comment{args, false}
}

// CommentSpew is like Commentf but preatty print all args using spew
func CommentSpew(args ...interface{}) CommentInterface {
	return comment{args, true}
}

func (c comment) CheckCommentString() string {
	if c.isSpew {
		args := make([]string, len(c.args), len(c.args))
		for i := range c.args {
			args[i] = spew.Sdump(c.args[i])
		}
		return strings.Join(args, " ")
	}
	return fmt.Sprint(c.args...)
}

// -----------------------------------------------------------------------
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

// Check if two numbers are close enough
var EqualsWithTolerance Checker = &equalsWithToleranceChecker{&CheckerInfo{Name: "EqualsWithTolerance", Params: []string{"obtained", "expected", "tolerance"}}}

// -----------------------------------------------------------------------
type betweenChecker struct {
	*CheckerInfo
}

// Between check if a numeric values is between two other values
var Between Checker = &betweenChecker{&CheckerInfo{Name: "Between", Params: []string{"obtained", "lower", "upper"}}}

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

// -----------------------------------------------------------------------
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

type isEmptyChecker struct {
	*CheckerInfo
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

// Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or a slice with len == 0.
// For example:
//
// c.Assert(v, IsEmpty)
var IsEmpty Checker = &isEmptyChecker{
	&CheckerInfo{Name: "IsEmpty", Params: []string{"value"}},
}

// -----------------------------------------------------------------------

type satisfiesChecker struct {
	*CheckerInfo
}

// Satisfies checks whether a value causes the argument
// function to return true. The function must be of
// type func(T) bool where the value being checked
// is assignable to T.
var Satisfies Checker = &satisfiesChecker{
	&CheckerInfo{
		Name:   "Satisfies",
		Params: []string{"obtained", "func(T) bool"},
	},
}

func (checker *satisfiesChecker) Check(params []interface{}, names []string) (result bool, error string) {
	f := reflect.ValueOf(params[1])
	ft := f.Type()
	if ft.Kind() != reflect.Func ||
		ft.NumIn() != 1 ||
		ft.NumOut() != 1 ||
		ft.Out(0) != reflect.TypeOf(true) {
		return false, fmt.Sprintf("expected func(T) bool, got %s", ft)
	}
	v := reflect.ValueOf(params[0])
	if !v.IsValid() {
		if !canBeNil(ft.In(0)) {
			return false, fmt.Sprintf("cannot assign nil to argument %T", ft.In(0))
		}
		v = reflect.Zero(ft.In(0))
	}
	if !v.Type().AssignableTo(ft.In(0)) {
		return false, fmt.Sprintf("wrong argument type %s for %s", v.Type(), ft)
	}
	return f.Call([]reflect.Value{v})[0].Interface().(bool), ""
}
func canBeNil(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice:
		return true
	}
	return false
}

// -----------------------------------------------------------------------
type hasPrefixChecker struct {
	*CheckerInfo
}

// HasPrefix checker for checking strings
var HasPrefix Checker = &hasPrefixChecker{
	&CheckerInfo{Name: "HasPrefix", Params: []string{"obtained", "expected"}},
}

func (checker *hasPrefixChecker) Check(params []interface{}, names []string) (result bool, error string) {
	expected, ok := params[1].(string)
	if !ok {
		return false, "expected must be a string"
	}

	obtained, isString := stringOrStringer(params[0])
	if isString {
		return strings.HasPrefix(obtained, expected), ""
	}

	return false, "Obtained value is not a string and has no .String()"
}

// -----------------------------------------------------------------------
type hasSuffixChecker struct {
	*CheckerInfo
}

// HasSuffix Checker
var HasSuffix Checker = &hasSuffixChecker{
	&CheckerInfo{Name: "HasSuffix", Params: []string{"obtained", "expected"}},
}

func (checker *hasSuffixChecker) Check(params []interface{}, names []string) (result bool, error string) {
	expected, ok := params[1].(string)
	if !ok {
		return false, "expected must be a string"
	}

	obtained, isString := stringOrStringer(params[0])
	if isString {
		return strings.HasSuffix(obtained, expected), ""
	}

	return false, "Obtained value is not a string and has no .String()"
}
