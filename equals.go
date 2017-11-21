package checkers

import (
	"fmt"
	"reflect"

	gc "gopkg.in/check.v1"
)

// -----------------------------------------------------------------------
type strEquals struct {
	*gc.CheckerInfo
}

func (c *strEquals) Check(params []interface{}, names []string) (result bool, error string) {
	p1, p2 := params[0], params[1]
	p1Nil, p2Nil := isNil(p1), isNil(p2)
	if p1Nil || p2Nil {
		if p1Nil == p2Nil {
			return true, ""
		}
		if p2Nil {
			return false, "expecting nil, got not-nil value"
		}
		return false, "expecting not-nil value, got nil"
	}

	return fmt.Sprint(p1) == fmt.Sprint(p2), ""
}

// StrEquals checks if fmt.Sprint values of objects are equal
var StrEquals gc.Checker = &strEquals{
	&gc.CheckerInfo{Name: "StrEquals", Params: []string{"obtained", "expected"}}}

// copied from gopkg.in/check.v1
func isNil(obtained interface{}) (result bool) {
	if obtained == nil {
		result = true
	} else {
		switch v := reflect.ValueOf(obtained); v.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			return v.IsNil()
		}
	}
	return
}
