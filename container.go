package checkers

import (
	"fmt"
	"reflect"

	gc "gopkg.in/check.v1"
	"strings"
)

type containsChecker struct {
	*gc.CheckerInfo
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

// Contains checker checks if an array, slice or string contains an element
var Contains gc.Checker = &containsChecker{
	&gc.CheckerInfo{Name: "Contains", Params: []string{"Container", "Value expected to contain"}}}

// -----------------------------------------------------------------------
type sliceEquals struct {
	*gc.CheckerInfo
}

func (c *sliceEquals) Check(params []interface{}, names []string) (result bool, error string) {
	s1 := params[0]
	s2 := params[1]
	vs1 := reflect.ValueOf(s1)
	vs2 := reflect.ValueOf(s2)

	if vs1.Kind() != reflect.Slice || vs2.Kind() != reflect.Slice {
		return false, "Both arguments must be slices"
	}
	l := vs1.Len()
	if l != vs2.Len() {
		return false, ""
	}
	return reflect.DeepEqual(s1, s2), ""
}

// SliceEquals check if two slices has the same values
var SliceEquals gc.Checker = &sliceEquals{
	&gc.CheckerInfo{Name: "SliceEquals", Params: []string{"obtained", "expected"}}}

// -----------------------------------------------------------------------
type mapEquals struct {
	*gc.CheckerInfo
}

func (c *mapEquals) Check(params []interface{}, names []string) (result bool, error string) {
	s1 := params[0]
	s2 := params[1]
	vs1 := reflect.ValueOf(s1)
	vs2 := reflect.ValueOf(s2)

	if vs1.Kind() != reflect.Map || vs2.Kind() != reflect.Map {
		return false, "Both arguments must be maps"
	}
	l := vs1.Len()
	if l != vs2.Len() {
		return false, ""
	}
	return reflect.DeepEqual(s1, s2), ""
}

// MapEquals check if two maps has the same values
var MapEquals gc.Checker = &mapEquals{
	&gc.CheckerInfo{Name: "MapEquals", Params: []string{"obtained", "expected"}}}

// -----------------------------------------------------------------------

type sameContents struct {
	*gc.CheckerInfo
}

// SameContents checks that the obtained slice contains all the values (and
// same number of values) of the expected slice and vice versa, without respect
// to order or duplicates. Uses DeepEquals on mapped contents to compare.
var SameContents gc.Checker = &sameContents{
	&gc.CheckerInfo{Name: "SameContents", Params: []string{"obtained", "expected"}},
}

func (checker *sameContents) Check(params []interface{}, names []string) (result bool, error string) {
	if len(params) != 2 {
		return false, "SameContents expects two slice arguments"
	}
	obtained := params[0]
	expected := params[1]

	tob := reflect.TypeOf(obtained)
	if tob.Kind() != reflect.Slice {
		return false, fmt.Sprintf("SameContents expects the obtained value to be a slice, got %q",
			tob.Kind())
	}

	texp := reflect.TypeOf(expected)
	if texp.Kind() != reflect.Slice {
		return false, fmt.Sprintf("SameContents expects the expected value to be a slice, got %q",
			texp.Kind())
	}

	if texp != tob {
		return false, fmt.Sprintf(
			"SameContents expects two slices of the same type, expected: %q, got: %q",
			texp, tob)
	}

	vexp := reflect.ValueOf(expected)
	vob := reflect.ValueOf(obtained)
	length := vexp.Len()

	if vob.Len() != length {
		// Slice has incorrect number of elements
		return false, ""
	}

	// spin up maps with the entries as keys and the counts as values
	mob := make(map[interface{}]int, length)
	mexp := make(map[interface{}]int, length)

	for i := 0; i < length; i++ {
		mexp[vexp.Index(i).Interface()]++
		mob[vob.Index(i).Interface()]++
	}
	return reflect.DeepEqual(mob, mexp), ""
}

// -----------------------------------------------------------------------

// -----------------------------------------------------------------------
