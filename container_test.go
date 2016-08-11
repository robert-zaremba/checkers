package checkers

import (
	"sort"

	. "gopkg.in/check.v1"
)

type ContainerSuite struct{}

func (s *ContainerSuite) TestContains(c *C) {
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

func (s *ContainerSuite) TestIsIn(c *C) {
	a := []int{2, 3, 4}
	c.Check(a[0], IsIn, a)
	c.Check(a[1], IsIn, a)
	c.Check(a[2], IsIn, a)
	c.Check(2, IsIn, a)
	c.Check(3, IsIn, a)
	c.Check(4, IsIn, a)
	c.Check(5, Not(IsIn), a)
	c.Check(a, Not(IsIn), a)
	c.Check("x", Not(IsIn), a)

	b := []x{x{"1"}, x{"2"}}
	c.Check(b[0], IsIn, b)
	c.Check(b[1], IsIn, b)
	c.Check(x{"1"}, IsIn, b)
	c.Check(x{"2"}, IsIn, b)
	c.Check(x{"3"}, Not(IsIn), b)
	c.Check(y{0}, Not(IsIn), b)

	c.Check("23", IsIn, "1234")
	c.Check("4", IsIn, "1234")
	c.Check("", IsIn, "1234")
	c.Check("0", Not(IsIn), "1234")
}

func (s *ContainerSuite) TestSliceEquals(c *C) {
	type Point struct{ X, Y int }
	c.Check([]string{}, SliceEquals, []string{})
	c.Check([]Point{}, SliceEquals, []Point{})
	c.Check([]Point{{1, 3}, {2, 10}}, SliceEquals, []Point{{1, 3}, {2, 10}})

	c.Check([]string{}, Not(SliceEquals), []int{})
	c.Check([]int{}, Not(SliceEquals), []int64{})
	c.Check([]int{1, 2}, Not(SliceEquals), []int64{2, 1})
	c.Check([]int{1, 2}, Not(SliceEquals), []int64{1, 2, 3})
	c.Check([]int{1, 2, 3}, Not(SliceEquals), []int64{1})
}

func (s *ContainerSuite) TestMapEquals(c *C) {
	type Point struct{ X, Y int }
	c.Check(map[int]string{}, MapEquals, map[int]string{})
	c.Check(map[int]Point{}, MapEquals, map[int]Point{})
	c.Check(map[int]Point{1: {1, 2}}, MapEquals, map[int]Point{1: {1, 2}})

	c.Check(map[int]string{}, Not(MapEquals), map[int]int{})
	c.Check(map[int]Point{1: {1, 2}}, Not(MapEquals), map[int]Point{2: {1, 2}})
	c.Check(map[int]Point{1: {1, 2}}, Not(MapEquals), map[int]Point{1: {2, 2}})
	c.Check(map[int]Point{1: {1, 2}}, Not(MapEquals), map[int]Point{1: {1, 2}, 2: {1, 2}})
	c.Check(map[int]Point{1: {1, 2}, 2: {1, 2}}, Not(MapEquals), map[int]Point{1: {1, 2}})
}

func (s *ContainerSuite) TestSameContents(c *C) {
	//// positive cases ////

	// same
	c.Check(
		[]int{1, 2, 3}, SameContents,
		[]int{1, 2, 3})

	// empty
	c.Check(
		[]int{}, SameContents,
		[]int{})

	// single
	c.Check(
		[]int{1}, SameContents,
		[]int{1})

	// different order
	c.Check(
		[]int{1, 2, 3}, SameContents,
		[]int{3, 2, 1})

	// multiple copies of same
	c.Check(
		[]int{1, 1, 2}, SameContents,
		[]int{2, 1, 1})

	type test struct {
		s string
		i int
	}

	// test structs
	c.Check(
		[]test{{"a", 1}, {"b", 2}}, SameContents,
		[]test{{"b", 2}, {"a", 1}})

	//// negative cases ////

	// different contents
	c.Check(
		[]int{1, 3, 2, 5}, Not(SameContents),
		[]int{5, 2, 3, 4})

	// different size slices
	c.Check(
		[]int{1, 2, 3}, Not(SameContents),
		[]int{1, 2})

	// different counts of same items
	c.Check(
		[]int{1, 1, 2}, Not(SameContents),
		[]int{1, 2, 2})

	/// Error cases ///
	//  note: for these tests, we can't use Not, since Not passes the error value through
	// and checks with a non-empty error always count as failed
	// Oddly, there doesn't seem to actually be a way to check for an error from a Checker.

	// different type
	res, err := SameContents.Check([]interface{}{
		[]string{"1", "2"},
		[]int{1, 2},
	}, []string{})
	c.Check(res, IsFalse)
	c.Check(err, Not(Equals), "")

	// obtained not a slice
	res, err = SameContents.Check([]interface{}{
		"test",
		[]int{1},
	}, []string{})
	c.Check(res, IsFalse)
	c.Check(err, Not(Equals), "")

	// expected not a slice
	res, err = SameContents.Check([]interface{}{
		[]int{1},
		"test",
	}, []string{})
	c.Check(res, IsFalse)
	c.Check(err, Not(Equals), "")
}

func (s *ContainerSuite) TestIsSorted(c *C) {
	var a sort.IntSlice
	c.Check(a, IsSorted)

	a = []int{2, 3, 4}
	c.Check(a, IsSorted)

	a = []int{-2, 0, 4}
	c.Check(a, IsSorted)

}
