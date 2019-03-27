/*
Package checkers is an extension to github.com/go-check/check library.
It provides additional usefull checkers:

  * Between - checks if a number is between given 2 other numbers
  * Contains (checks if a slice/array/string contains specified element)
  * CloseTo - an alias for EqualsWithTolerance
  * IsIn (checks if an element is in a slice/array/string)
  * DoesNotExist - checks if a path exists
  * EqualsWithTolerance - checks if two numbers are "close enough"
  * HasPrefix, HasSuffix
  * IsDirectory
  * IsEmpty - checks if specified object is empty (nil, [], {}, "", 0)
  * IsNonEmptyFile
  * IsSymlink, SymlinkDoesNotExist
  * IsTrue, IsFalse
  * MapEquals - checks if 2 maps contain the same elements
  * SameContent - multiset compairson. Checks if two slices contain same elements (including duplicates), ignoring the order.
  * SamePath - follows OS symlink to check if two paths are same.
  * Satisfies - check if a value satisfies functional predicate
  * SliceEquals - checks if 2 slices contain the same elements
  * StrEquals - checks if fmt.Sprint values of objects are equal
  * TimeEquals - checks if time is the same up to microseconds, useful if some driver or type truncates the nanosecond time accuracy.
  * WithinDuration - checks if an obtained time is not earlier/later than the expected time + duration
  * DurationLessThan

Furthermore there are two additional CommentInterface implementations:

  * Comment - works like fmt.Sprint (doesn't have a formatting string)
  * CommentSpew - like Commentf, but uses Spew to format the structures.
*/
package checkers
