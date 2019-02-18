package checkers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
	gc "gopkg.in/check.v1"
)

type comment struct {
	args   []interface{}
	isSpew bool
}

// Comment is like Commentf but without formatting string
func Comment(args ...interface{}) gc.CommentInterface {
	return comment{args, false}
}

// CommentSpew is like Commentf but preatty print all args using spew
func CommentSpew(args ...interface{}) gc.CommentInterface {
	return comment{args, true}
}

func (c comment) CheckCommentString() string {
	if c.isSpew {
		args := make([]string, len(c.args))
		for i := range c.args {
			args[i] = spew.Sdump(c.args[i])
		}
		return strings.Join(args, " ")
	}
	return fmt.Sprint(c.args...)
}

// -----------------------------------------------------------------------
type isTrueChecker struct {
	*gc.CheckerInfo
}

// IsTrue checks if value == true
// For example:
//
//	c.Assert(v, IsTrue)
var IsTrue gc.Checker = &isTrueChecker{
	&gc.CheckerInfo{Name: "IsTrue", Params: []string{"value"}},
}

func (checker *isTrueChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0] == true, ""
}

// -----------------------------------------------------------------------

type isFalseChecker struct {
	*gc.CheckerInfo
}

// IsFalse checks if value == false
// For example:
//
//	c.Assert(v, IsFalse)
var IsFalse gc.Checker = &isFalseChecker{
	&gc.CheckerInfo{Name: "IsFalse", Params: []string{"value"}},
}

func (checker *isFalseChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0] == false, ""
}

// -----------------------------------------------------------------------

type isEmptyChecker struct {
	*gc.CheckerInfo
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
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return objValue.Len() == 0, ""
	case reflect.Ptr:
		{
			if objValue.IsNil() {
				return
			}
		}
	}

	return false, ""
}

// IsEmpty asserts that the specified object is empty. I.e. nil, "", false, 0 or a slice with len == 0.
// For example:
//
// c.Assert(v, IsEmpty)
var IsEmpty gc.Checker = &isEmptyChecker{
	&gc.CheckerInfo{Name: "IsEmpty", Params: []string{"value"}},
}

// -----------------------------------------------------------------------

type satisfiesChecker struct {
	*gc.CheckerInfo
}

// Satisfies checks whether a value causes the argument
// function to return true. The function must be of
// type func(T) bool where the value being checked
// is assignable to T.
var Satisfies gc.Checker = &satisfiesChecker{
	&gc.CheckerInfo{
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
	*gc.CheckerInfo
}

// HasPrefix checker for checking strings
var HasPrefix gc.Checker = &hasPrefixChecker{
	&gc.CheckerInfo{Name: "HasPrefix", Params: []string{"obtained", "expected"}},
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
	*gc.CheckerInfo
}

// HasSuffix Checker
var HasSuffix gc.Checker = &hasSuffixChecker{
	&gc.CheckerInfo{Name: "HasSuffix", Params: []string{"obtained", "expected"}},
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

// -----------------------------------------------------------------------
type errorContains struct {
	*gc.CheckerInfo
}

// ErrorContains checks if the error message (output of the .String() method)
// contains the given string.
var ErrorContains gc.Checker = &errorContains{
	&gc.CheckerInfo{Name: "ErrorContains", Params: []string{"obtained", "expected"}},
}

type errorI interface {
	Error() string
}

func (checker *errorContains) Check(params []interface{}, names []string) (result bool, error string) {
	expected, ok := params[1].(string)
	if !ok {
		return false, "expected must be a string"
	}

	if params[0] == nil {
		return false, "obtained nil"
	}
	err, isErr := params[0].(errorI)
	if !isErr {
		return false, "Obtained value doesn't implement error interface"
	}
	msg := err.Error()
	return strings.Contains(msg, expected), ""
}
